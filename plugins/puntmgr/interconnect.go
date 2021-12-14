// SPDX-License-Identifier: Apache-2.0

// Copyright 2021 PANTHEON.tech
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package puntmgr

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/golang/protobuf/proto"

	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/cn-infra/v2/servicelabel"

	"go.ligato.io/vpp-agent/v3/client"
	"go.ligato.io/vpp-agent/v3/plugins/linux/nsplugin"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
	linux_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/linux/interfaces"
	linux_namespace "go.ligato.io/vpp-agent/v3/proto/ligato/linux/namespace"
	vpp_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	vpp_l3 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l3"

	pb "go.pantheon.tech/stonework/proto/puntmgr"
)

const (
	memifSockDir = "/run/stonework/memif"
)

// InterconnectManager manages creation/deletion and sharing of VPP<->CNF/Linux interconnects.
type InterconnectManager interface {
	// Add new VPP<->CNF/Linux interconnects needed for a given punt.
	// localTxn = configuration items to configure on this side (StoneWork / Standalone CNF)
	// remoteTxn = configuration items to configure on the side of the StoneWork module
	AddInterconnects(localTxn, remoteTxn client.ChangeRequest, puntId puntID, reqs []InterconnectReq,
		icType pb.PuntRequest_InterconnectType, withMultiplex bool) (interconnects []*pb.PuntMetadata_Interconnect, err error)
	// Delete all VPP<->CNF/Linux interconnects created for a given punt.
	DelInterconnects(localTxn, remoteTxn client.ChangeRequest, puntId puntID) (err error)
	// GetLinuxVrfName returns the name used for Linux VRF device corresponding to the given VPP VRF.
	// Method is "static" in the sense that it can be called anytime, regardless of the internal state of the Manager.
	GetLinuxVrfName(vrf uint32) string
}

// interconnectManager implements InterconnectManager interface.
type interconnectManager struct {
	log      logging.Logger
	ifPlugin ifplugin.API
	svcLabel servicelabel.ReaderAPI
	netNsReg NetNsRegistry

	allocCidr       *net.IPNet
	nextAllocSubnet int

	icByID          map[icID]*interconnect
	icByVppSelector map[string][]*interconnect // key = vpp selector
	icByPuntID      map[puntID][]*interconnect

	proxiedIfaces map[string]*proxiedIface // key = unnumberedToIface

	vrfRefCount map[vrfID][]*vrfRefCountPerCnf
}

// Unlike pb.PuntMetadata_InterconnectID this can be also used as a map key.
type icID struct {
	// What/where packets are punted on the VPP side using this interconnect.
	// Each punt handler (there is one for each PuntRequest.Type) defines its own selectors.
	VppSelector string
	// What/where packets are punted on the CNF side using this interconnect.
	// Generated by PuntManager. Outside manager only useful in combination with vpp_selector to obtain
	// unique id for the interconnect.
	CnfSelector string
}

func (id icID) String() string {
	return id.VppSelector + "|" + id.CnfSelector
}

// vrfID identifies VPP VRF replicated in the given Linux/CNF network stack (e.g. using Linux VRF device).
type vrfID struct {
	vppVrf      uint32
	CnfSelector string
}

func (id vrfID) String() string {
	return fmt.Sprintf("%d|%s", id.vppVrf, id.CnfSelector)
}

// This structure takes count of all references to a given VRF from a given CNF.
type vrfRefCountPerCnf struct {
	cnfMsLabel string
	refCount   int
}

func (rc vrfRefCountPerCnf) String() string {
	return fmt.Sprintf("%s->%d", rc.cnfMsLabel, rc.refCount)
}

// VPP <-> CNF/Linux interconnect.
type interconnect struct {
	id             icID
	request        InterconnectReq
	icType         pb.PuntRequest_InterconnectType
	withMultiplex  bool
	metadata       *pb.PuntMetadata_Interconnect
	proxyIfaceName string
	allocSubnetIdx int
	usedBy         []puntID
}

// Interface that is effectively proxied (from the point of view of a CNF) using an unnumbered
// VPP interface and a proxy ARP.
type proxiedIface struct {
	name      string
	ips       []*net.IPNet
	proxiedBy []icID
}

func NewInterconnectManager(log logging.Logger, ifPlugin ifplugin.API, svcLabel servicelabel.ReaderAPI, nsPlugin nsplugin.API,
	allocCidr *net.IPNet) InterconnectManager {
	_ = os.Mkdir(memifSockDir, os.ModeDir)
	return &interconnectManager{
		log:             log,
		ifPlugin:        ifPlugin,
		svcLabel:        svcLabel,
		allocCidr:       allocCidr,
		netNsReg:        NewNetNsRegistry(nsPlugin, svcLabel),
		icByID:          make(map[icID]*interconnect),
		icByVppSelector: make(map[string][]*interconnect),
		icByPuntID:      make(map[puntID][]*interconnect),
		proxiedIfaces:   make(map[string]*proxiedIface),
		vrfRefCount:     make(map[vrfID][]*vrfRefCountPerCnf),
	}
}

// Add new VPP<->CNF/Linux interconnects needed for a given punt.
func (m *interconnectManager) AddInterconnects(localTxn, remoteTxn client.ChangeRequest, puntId puntID, reqs []InterconnectReq,
	icType pb.PuntRequest_InterconnectType, withMultiplex bool) (resp []*pb.PuntMetadata_Interconnect, err error) {

	if _, hasICs := m.icByPuntID[puntId]; hasICs {
		return nil, fmt.Errorf("punt %v already has some interconnects created", puntId)
	}

	var ics []*interconnect
	proxiedIfaces := make(map[string]*proxiedIface)
	vppSelectors := make(map[string]struct{})
	nextAllocSubnet := m.nextAllocSubnet

	// 1. build definition of each interconnect and check for conflicts
	//    without making any changes to any of the internal maps
	for _, req := range reqs {
		// build interconnect ID
		cnfSelector, err := m.getCnfSelector(puntId, req, icType)
		if err != nil {
			return nil, err
		}
		id := icID{
			VppSelector: req.vppSelector,
			CnfSelector: cnfSelector,
		}
		m.log.Infof("Configuring interconnect with ID %s", id.String())

		// check for conflicts
		if _, duplicate := vppSelectors[req.vppSelector]; duplicate {
			return nil, fmt.Errorf("duplicate VPP selector %s", req.vppSelector)
		}
		allocSubnetIdx := nextAllocSubnet
		vppSelectors[req.vppSelector] = struct{}{}
		for _, ic2 := range m.icByVppSelector[req.vppSelector] {
			if !withMultiplex || !ic2.withMultiplex {
				return nil, fmt.Errorf("VPP selector %s is busy", req.vppSelector)
			}
			if ic2.metadata.Id.CnfSelector == cnfSelector {
				if ic2.icType != pb.PuntRequest_TAP || icType != pb.PuntRequest_TAP ||
					!ic2.request.link.equivalent(req.link) {
					return nil, fmt.Errorf("CNF selector %s is busy", cnfSelector)
				}
				// it will be shared
				allocSubnetIdx = ic2.allocSubnetIdx
			}
		}

		// handle proxied interface
		var (
			proxyIface     *proxiedIface
			proxyIfaceName string
		)
		if ifLink, isIfLink := req.link.(*InterfaceLink); isIfLink {
			if proxyIfaceName = ifLink.unnumberedToIface; proxyIfaceName != "" {
				_, duplicate := proxiedIfaces[proxyIfaceName]
				if duplicate {
					return nil, fmt.Errorf("duplicate request to proxy VPP interface %s", proxyIfaceName)
				}
				if proxyIface2, alreadyProxied := m.proxiedIfaces[proxyIfaceName]; alreadyProxied {
					proxyIface = &proxiedIface{
						name:      proxyIfaceName,
						ips:       proxyIface2.ips, // no need to copy, never changed
						proxiedBy: make([]icID, len(proxyIface2.proxiedBy)+1),
					}
					copy(proxyIface.proxiedBy, proxyIface2.proxiedBy)
					proxyIface.proxiedBy[len(proxyIface.proxiedBy)-1] = id
				} else {
					proxyIface, err = m.buildProxyInterface(proxyIfaceName, id)
					m.log.Debugf("buildProxyInterface: %v", proxyIface)
					if err != nil {
						return nil, err
					}
				}
				proxiedIfaces[proxyIface.name] = proxyIface
			}
		}

		// allocate subnet if requested
		var allocdSubnet *net.IPNet // nil or exactly two host IPs
		if ifLink, isIfLink := req.link.(*InterfaceLink); isIfLink {
			if ifLink.allocateSubnet {
				ones, bits := m.allocCidr.Mask.Size()
				if bits-ones < 2 {
					return nil, errors.New("failed to allocate subnet for interconnect: " +
						"cidr for address allocation is too small")
				}
				allocdSubnet, err = cidr.Subnet(m.allocCidr, bits-ones-2, allocSubnetIdx)
				if err != nil {
					return nil, fmt.Errorf("failed to allocate subnet for interconnect: %v", err)
				}
				if allocSubnetIdx == nextAllocSubnet {
					nextAllocSubnet++
				}
			}
		}

		// build interconnect definition
		metadata := m.buildMetadata(puntId, id, icType, req.link, allocdSubnet, proxyIface)
		m.log.Debugf("Interconnect metadata: %+v", metadata)
		ics = append(ics, &interconnect{
			id:             id,
			request:        req,
			icType:         icType,
			withMultiplex:  withMultiplex,
			metadata:       metadata,
			proxyIfaceName: proxyIfaceName,
			allocSubnetIdx: allocSubnetIdx,
			usedBy:         []puntID{puntId},
		})
	}

	// nothing can fail from this point on...
	m.nextAllocSubnet = nextAllocSubnet

	// update internal maps and prepare transaction
	for _, ic := range ics {
		sameIC, sharedIC := m.icByID[ic.id]
		var (
			icSharedWithin  bool
			vrfSharedWithin bool
			sharedVrf       bool
		)
		// -> Interconnect
		if sharedIC {
			m.log.Debugf("Interconnect %s is shared", ic.id)
			for _, usedBy := range sameIC.usedBy {
				if usedBy.cnfMsLabel == puntId.cnfMsLabel {
					m.log.Debugf("Interconnect %s is shared within the same CNF %s",
						ic.id, usedBy.cnfMsLabel)
					icSharedWithin = true
				}
			}
			sameIC.metadata.Shared = true
			sameIC.usedBy = append(sameIC.usedBy, puntId)
			m.icByPuntID[puntId] = append(m.icByPuntID[puntId], sameIC)
			resp = append(resp, sameIC.metadata)
			// verify that metadata are built deterministically
			ic.metadata.Shared = true
			if !proto.Equal(ic.metadata, sameIC.metadata) {
				panic(fmt.Sprintf("metadata generated for the same IC do not match (%+v vs. %+v)",
					ic.metadata, sameIC.metadata))
			}
		} else {
			// new IC
			m.icByID[ic.id] = ic
			m.icByVppSelector[ic.id.VppSelector] = append(m.icByVppSelector[ic.id.VppSelector], ic)
			m.icByPuntID[puntId] = append(m.icByPuntID[puntId], ic)
			resp = append(resp, ic.metadata)
		}
		if !icSharedWithin {
			m.buildInterconnectTxn(localTxn, remoteTxn, ic, sharedIC, false)
		}
		// -> VRF
		if ic.metadata.CnfInterface != nil {
			vrf := ic.metadata.CnfInterface.VrfRT
			vrfID := vrfID{vppVrf: vrf, CnfSelector: ic.id.CnfSelector}
			var counted bool
			for _, rc := range m.vrfRefCount[vrfID] {
				if rc.refCount > 0 {
					sharedVrf = true
				}
				if rc.cnfMsLabel == puntId.cnfMsLabel {
					rc.refCount += 1
					vrfSharedWithin = rc.refCount > 1
					counted = true
				}
			}
			if !counted {
				m.vrfRefCount[vrfID] = append(m.vrfRefCount[vrfID], &vrfRefCountPerCnf{
					cnfMsLabel: puntId.cnfMsLabel,
					refCount:   1,
				})
			}
			m.log.Debugf("vrfRefCount=%v after adding ic=%v for CNF=%s",
				m.vrfRefCount, ic.id, puntId.cnfMsLabel)
		}
		if !vrfSharedWithin {
			m.buildVrfTxn(localTxn, remoteTxn, ic, sharedVrf, false)
		}
	}

	// rebuild proxy ARP (the only global configuration item)
	if len(proxiedIfaces) > 0 {
		var newProxyIface bool
		for _, proxyIface := range proxiedIfaces {
			if _, shared := m.proxiedIfaces[proxyIface.name]; !shared {
				newProxyIface = true
			}
			m.proxiedIfaces[proxyIface.name] = proxyIface
		}
		if newProxyIface {
			// rebuild proxy ARP
			m.rebuildProxyArp(localTxn)
		}
	}
	return resp, nil
}

// Delete all VPP<->CNF/Linux interconnects created for a given punt.
func (m *interconnectManager) DelInterconnects(localTxn, remoteTxn client.ChangeRequest, puntId puntID) (err error) {
	ics, hasICs := m.icByPuntID[puntId]
	if !hasICs {
		return fmt.Errorf("no interconnects are configured for punt %v", puntId)
	}

	// nothing can fail from this point on...

	var updateProxyArp bool
	for _, ic := range ics {
		sharedIC := len(ic.usedBy) > 1
		var (
			icSharedWithin  bool
			vrfSharedWithin bool
			sharedVrf       bool
		)
		// -> Interconnect
		if sharedIC {
			// remove only items configured on the CNF side
			var sharedWith []puntID
			for _, ic2 := range ic.usedBy {
				if ic2 != puntId {
					sharedWith = append(sharedWith, ic2)
					if ic2.cnfMsLabel == puntId.cnfMsLabel {
						icSharedWithin = true
					}
				}
			}
			ic.usedBy = sharedWith
			ic.metadata.Shared = len(ic.usedBy) > 1
		} else {
			// remove entire IC
			if ic.proxyIfaceName != "" {
				proxyIface := m.proxiedIfaces[ic.proxyIfaceName]
				var icWithSameProxyIf []icID
				for _, ic2 := range proxyIface.proxiedBy {
					if ic2 != ic.id {
						icWithSameProxyIf = append(icWithSameProxyIf, ic2)
					}
				}
				if len(icWithSameProxyIf) > 0 {
					proxyIface.proxiedBy = icWithSameProxyIf
				} else {
					delete(m.proxiedIfaces, ic.proxyIfaceName)
					updateProxyArp = true
				}
			}
			var icWithSameVppSel []*interconnect
			for _, ic2 := range m.icByVppSelector[ic.id.VppSelector] {
				if ic2.id != ic.id {
					icWithSameVppSel = append(icWithSameVppSel, ic2)
				}
			}
			if len(icWithSameVppSel) > 0 {
				m.icByVppSelector[ic.id.VppSelector] = icWithSameVppSel
			} else {
				delete(m.icByVppSelector, ic.id.VppSelector)
			}
			delete(m.icByID, ic.id)
		}
		if !icSharedWithin {
			m.buildInterconnectTxn(localTxn, remoteTxn, ic, sharedIC, true)
		}
		// -> VRF
		if ic.metadata.CnfInterface != nil {
			vrf := ic.metadata.CnfInterface.VrfRT
			vrfID := vrfID{vppVrf: vrf, CnfSelector: ic.id.CnfSelector}
			for _, rc := range m.vrfRefCount[vrfID] {
				if rc.cnfMsLabel == puntId.cnfMsLabel {
					rc.refCount -= 1
					vrfSharedWithin = rc.refCount > 0
				}
				if rc.refCount > 0 {
					sharedVrf = true
				}
			}
			m.log.Debugf("vrfRefCount=%v after removing ic=%v for CNF=%s",
				m.vrfRefCount, ic.id, puntId.cnfMsLabel)
		}
		if !vrfSharedWithin {
			m.buildVrfTxn(localTxn, remoteTxn, ic, sharedVrf, true)
		}
	}
	if updateProxyArp {
		m.rebuildProxyArp(localTxn)
	}
	delete(m.icByPuntID, puntId)
	return nil
}

func (m *interconnectManager) getCnfSelector(
	puntId puntID, icReq InterconnectReq, icType pb.PuntRequest_InterconnectType) (string, error) {
	switch link := icReq.link.(type) {
	case *AFUnixLink:
		if icType != pb.PuntRequest_AF_UNIX {
			return "", errors.New("interconnect link/type mismatch")
		}
		return "af-unix::" + link.socketPath, nil
	case *InterfaceLink:
		switch icType {
		case pb.PuntRequest_MEMIF:
			return "memif::" + m.getMemifSuffix(puntId, icReq.vppSelector), nil
		case pb.PuntRequest_TAP:
			nsId, err := m.netNsReg.GetNetNsID(puntId.cnfMsLabel)
			if err != nil {
				err = fmt.Errorf("failed to obtains net-ns ID for microservice %s: %w",
					puntId.cnfMsLabel, err)
				return "", err
			}
			return "netns::" + strconv.Itoa(nsId), nil
		case pb.PuntRequest_AF_UNIX:
			return "", errors.New("interconnect link/type mismatch")
		default:
			return "", errors.New("unrecognized interconnect type")
		}
	}
	return "", errors.New("unrecognized interconnect link")
}

func (m *interconnectManager) buildMetadata(
	puntId puntID, icID icID, icType pb.PuntRequest_InterconnectType, link InterconnectLink,
	allocdSubnet *net.IPNet, proxyIface *proxiedIface) *pb.PuntMetadata_Interconnect {
	var vppIface, cnfIface *pb.PuntMetadata_Interface
	if ifLink, isIfLink := link.(*InterfaceLink); isIfLink {
		vppIface = &pb.PuntMetadata_Interface{}
		cnfIface = &pb.PuntMetadata_Interface{}
		// Interface name
		ifaceName := m.getIcIfaceName(icID, icType)
		if ifLink.interfaceName != "" {
			vppIface.Name = ifLink.interfaceName
			cnfIface.Name = ifLink.interfaceName
		} else {
			vppIface.Name = ifaceName
			cnfIface.Name = ifaceName
		}
		// MAC
		if ifLink.physAddress != "" {
			vppIface.PhysAddress = ifLink.physAddress
		} else {
			vppIface.PhysAddress = m.getIcIfaceMAC(ifaceName, true)
		}
		cnfIface.PhysAddress = m.getIcIfaceMAC(ifaceName, false)
		// IP addresses
		if len(ifLink.ipAddresses) != 0 {
			vppIface.IpAddresses = ifLink.ipAddresses
		} else if ifLink.unnumberedToIface != "" {
			for _, ip := range proxyIface.ips {
				cnfIface.IpAddresses = append(cnfIface.IpAddresses, ip.String())
			}
		} else if allocdSubnet != nil {
			vppIP, err := cidr.Host(allocdSubnet, 1)
			if err != nil {
				// should be unreachable
				m.log.Error(err)
			}
			vppIPNet := &net.IPNet{IP: vppIP, Mask: allocdSubnet.Mask}
			cnfIP, err := cidr.Host(allocdSubnet, 2)
			if err != nil {
				// should be unreachable
				m.log.Error(err)
			}
			cnfIPNet := &net.IPNet{IP: cnfIP, Mask: allocdSubnet.Mask}
			vppIface.IpAddresses = []string{vppIPNet.String()}
			cnfIface.IpAddresses = []string{cnfIPNet.String()}
		}
		// VRF
		vppIface.VrfRT = ifLink.vrf
		if !ifLink.withoutCNFVrf {
			cnfIface.VrfRT = ifLink.vrf
			if ifLink.vrf != 0 {
				cnfIface.VrfName = m.GetLinuxVrfName(ifLink.vrf)
			}
		}
	}
	return &pb.PuntMetadata_Interconnect{
		Id: &pb.PuntMetadata_InterconnectID{
			VppSelector: icID.VppSelector,
			CnfSelector: icID.CnfSelector,
		},
		VppInterface: vppIface,
		CnfInterface: cnfIface,
		// .Shared is updated during merge
	}
}

// buildInterconnectTxn prepares items to configure locally as well as remotely in order to build VPP<->CNF interconnect.
func (m *interconnectManager) buildInterconnectTxn(localTxn, remoteTxn client.ChangeRequest, ic *interconnect, sharedIC, remove bool) {
	switch ic.icType {
	case pb.PuntRequest_AF_UNIX:
		// Nothing to configure between VPP and CNF
		return
	case pb.PuntRequest_TAP:
		// get the ms label designated to reference this network namespace
		nsID, err := m.netNsReg.GetNetNsID(ic.usedBy[0].cnfMsLabel)
		isLocalCnf := nsID == 0
		if err != nil {
			// this should be unreachable
			m.log.Error(err)
		}
		var msLabel string
		if !isLocalCnf {
			msLabel, err = m.netNsReg.GetNetNsLabel(nsID)
			if err != nil {
				// this should be unreachable
				m.log.Error(err)
			}
		}
		// handle unnumbered interface
		link := ic.request.link.(*InterfaceLink)
		var unnumbered *vpp_interfaces.Interface_Unnumbered
		if link.unnumberedToIface != "" {
			unnumbered = &vpp_interfaces.Interface_Unnumbered{
				InterfaceWithIp: link.unnumberedToIface,
			}
		}
		// VPP side of the interconnect
		vppIface := &vpp_interfaces.Interface{
			Name:          ic.metadata.VppInterface.Name,
			Type:          vpp_interfaces.Interface_TAP,
			Enabled:       true,
			PhysAddress:   ic.metadata.VppInterface.PhysAddress,
			IpAddresses:   ic.metadata.VppInterface.IpAddresses,
			Vrf:           ic.metadata.VppInterface.VrfRT,
			SetDhcpClient: link.withDhcpClient,
			Mtu:           link.mtu,
			Unnumbered:    unnumbered,
			Link: &vpp_interfaces.Interface_Tap{
				Tap: &vpp_interfaces.TapLink{
					Version:        2,
					ToMicroservice: msLabel,
				},
			},
		}
		if !sharedIC {
			if remove {
				localTxn.Delete(vppIface)
			} else {
				localTxn.Update(vppIface)
			}
		}
		// Linux side of the interconnect
		var linuxNs *linux_namespace.NetNamespace
		if !isLocalCnf {
			linuxNs = &linux_namespace.NetNamespace{
				Type:      linux_namespace.NetNamespace_MICROSERVICE,
				Reference: msLabel,
			}
		}
		linuxIface := &linux_interfaces.Interface{
			Name:        ic.metadata.CnfInterface.Name,
			Type:        linux_interfaces.Interface_TAP_TO_VPP,
			Namespace:   linuxNs,
			Enabled:     true,
			IpAddresses: ic.metadata.CnfInterface.IpAddresses,
			PhysAddress: ic.metadata.CnfInterface.PhysAddress,
			Mtu:         link.mtu,
			Link: &linux_interfaces.Interface_Tap{
				Tap: &linux_interfaces.TapLink{
					VppTapIfName: vppIface.Name,
				},
			},
			VrfMasterInterface: ic.metadata.CnfInterface.VrfName,
		}
		if !sharedIC {
			if remove {
				localTxn.Delete(linuxIface)
			} else {
				localTxn.Update(linuxIface)
			}
		}
		if !isLocalCnf {
			existingLinuxIface := &linux_interfaces.Interface{
				Name:               linuxIface.Name,
				Type:               linux_interfaces.Interface_EXISTING,
				Enabled:            true,
				IpAddresses:        linuxIface.IpAddresses,
				LinkOnly:           true, // wait for IP addresses, do not configure them
				VrfMasterInterface: linuxIface.VrfMasterInterface,
			}
			if remove {
				remoteTxn.Delete(existingLinuxIface)
			} else {
				remoteTxn.Update(existingLinuxIface)
			}
		}

	case pb.PuntRequest_MEMIF:
		memifSufix := m.getMemifSuffix(ic.usedBy[0], ic.metadata.Id.VppSelector)
		memifSockPath := path.Join(memifSockDir, "memif-"+memifSufix+".sock")
		// handle unnumbered interface
		link := ic.request.link.(*InterfaceLink)
		var unnumbered *vpp_interfaces.Interface_Unnumbered
		if link.unnumberedToIface != "" {
			unnumbered = &vpp_interfaces.Interface_Unnumbered{
				InterfaceWithIp: link.unnumberedToIface,
			}
		}
		// VPP side of the interconnect
		vppIface := &vpp_interfaces.Interface{
			Name:          ic.metadata.VppInterface.Name,
			Type:          vpp_interfaces.Interface_MEMIF,
			Enabled:       true,
			PhysAddress:   ic.metadata.VppInterface.PhysAddress,
			IpAddresses:   ic.metadata.VppInterface.IpAddresses,
			Vrf:           ic.metadata.VppInterface.VrfRT,
			SetDhcpClient: link.withDhcpClient,
			Mtu:           link.mtu,
			Unnumbered:    unnumbered,
			Link: &vpp_interfaces.Interface_Memif{
				Memif: &vpp_interfaces.MemifLink{
					Mode:           vpp_interfaces.MemifLink_ETHERNET,
					Master:         true,
					Id:             1,
					SocketFilename: memifSockPath,
					Secret:         memifSufix,
				},
			},
		}
		if remove {
			localTxn.Delete(vppIface)
		} else {
			localTxn.Update(vppIface)
		}
		// CNF side of the interconnect
		cnfIface := &vpp_interfaces.Interface{
			Name:        ic.metadata.CnfInterface.Name,
			Type:        vpp_interfaces.Interface_MEMIF,
			Enabled:     true,
			PhysAddress: ic.metadata.CnfInterface.PhysAddress,
			IpAddresses: ic.metadata.CnfInterface.IpAddresses,
			Vrf:         ic.metadata.VppInterface.VrfRT,
			Mtu:         link.mtu,
			Link: &vpp_interfaces.Interface_Memif{
				Memif: &vpp_interfaces.MemifLink{
					Mode:           vpp_interfaces.MemifLink_ETHERNET,
					Master:         false,
					Id:             1,
					SocketFilename: memifSockPath,
					Secret:         memifSufix,
				},
			},
		}
		if remove {
			remoteTxn.Delete(cnfIface)
		} else {
			remoteTxn.Update(cnfIface)
		}
	}
	return
}

// buildVrfTxn prepares items to configure locally as well as remotely in order to replicate VPP VRFs in Linux.
func (m *interconnectManager) buildVrfTxn(localTxn, remoteTxn client.ChangeRequest, ic *interconnect, sharedVRF, remove bool) {
	switch ic.icType {
	case pb.PuntRequest_AF_UNIX:
		fallthrough
	case pb.PuntRequest_MEMIF:
		// Linux VRF devices not used (CNF may represent VRF inside in its own way)
		return
	case pb.PuntRequest_TAP:
		if ic.metadata.CnfInterface.VrfName == "" {
			// VPP VRF 0 = default routing table in Linux (which always exists)
			return
		}
		// get the ms label designated to reference this network namespace
		nsID, err := m.netNsReg.GetNetNsID(ic.usedBy[0].cnfMsLabel)
		isLocalCnf := nsID == 0
		if err != nil {
			// this should be unreachable
			m.log.Error(err)
		}
		var msLabel string
		if !isLocalCnf {
			msLabel, err = m.netNsReg.GetNetNsLabel(nsID)
			if err != nil {
				// this should be unreachable
				m.log.Error(err)
			}
		}
		var linuxNs *linux_namespace.NetNamespace
		if !isLocalCnf {
			linuxNs = &linux_namespace.NetNamespace{
				Type:      linux_namespace.NetNamespace_MICROSERVICE,
				Reference: msLabel,
			}
		}
		vrfDev := &linux_interfaces.Interface{
			Name:      ic.metadata.CnfInterface.VrfName,
			Type:      linux_interfaces.Interface_VRF_DEVICE,
			Namespace: linuxNs,
			Enabled:   true,
			Link: &linux_interfaces.Interface_VrfDev{
				VrfDev: &linux_interfaces.VrfDevLink{
					RoutingTable: ic.metadata.CnfInterface.VrfRT,
				},
			},
		}
		if !sharedVRF {
			if remove {
				localTxn.Delete(vrfDev)
			} else {
				localTxn.Update(vrfDev)
			}
		}
		if !isLocalCnf {
			existingLinuxVrf := &linux_interfaces.Interface{
				Name:     vrfDev.Name,
				Type:     linux_interfaces.Interface_EXISTING,
				Enabled:  true,
				LinkOnly: true,
			}
			if remove {
				remoteTxn.Delete(existingLinuxVrf)
			} else {
				remoteTxn.Update(existingLinuxVrf)
			}
		}
	}
	return
}

func (m *interconnectManager) rebuildProxyArp(localTxn client.ChangeRequest) {
	proxyArp := &vpp_l3.ProxyARP{}
	for _, proxyIface := range m.proxiedIfaces {
		icIfaces := make(map[string]struct{})
		vrfs := make(map[uint32]struct{})
		for _, icID := range proxyIface.proxiedBy {
			ic := m.icByID[icID]
			icIfaces[ic.metadata.VppInterface.Name] = struct{}{}
			vrfs[ic.metadata.VppInterface.VrfRT] = struct{}{}
		}
		for icIface := range icIfaces {
			proxyArp.Interfaces = append(proxyArp.Interfaces, &vpp_l3.ProxyARP_Interface{
				Name: icIface,
			})
		}
		for _, ipNet := range proxyIface.ips {
			if ipNet.IP.To4() == nil {
				continue
			}
			network := &net.IPNet{
				IP:   ipNet.IP.Mask(ipNet.Mask),
				Mask: ipNet.Mask,
			}
			firstIP, lastIP := cidr.AddressRange(network)
			for vrf := range vrfs {
				proxyArp.Ranges = append(proxyArp.Ranges, &vpp_l3.ProxyARP_Range{
					FirstIpAddr: firstIP.String(),
					LastIpAddr:  lastIP.String(),
					VrfId:       vrf,
				})
			}
		}

	}
	localTxn.Update(proxyArp)
}

func (m *interconnectManager) buildProxyInterface(name string, interconnectID icID) (*proxiedIface, error) {
	ifMeta, exists := m.ifPlugin.GetInterfaceIndex().LookupByName(name)
	if !exists || ifMeta == nil {
		return nil, fmt.Errorf(
			"VPP interface %s required for proxying was not found", name)
	}
	if len(ifMeta.IPAddresses) == 0 {
		return nil, fmt.Errorf(
			"VPP interface %s required for proxying does not have any IP address assigned", name)
	}
	proxyIface := &proxiedIface{
		name:      name,
		proxiedBy: []icID{interconnectID},
	}
	for _, ip := range ifMeta.IPAddresses {
		ipAddr, ipNet, err := net.ParseCIDR(ip)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IP address %s assigned to VPP interface %s: %w",
				ip, name, err)
		}
		if ipAddr.To4() != nil {
			ipAddr = ipAddr.To4()
		} else {
			ipAddr = ipAddr.To16()
		}
		ipNet.IP = ipAddr
		proxyIface.ips = append(proxyIface.ips, ipNet)
	}
	return proxyIface, nil
}

// getMemifSuffix returns suffix to use for memif socket, secret and also as a CNF selector
// for MEMIF interconnect.
func (m *interconnectManager) getMemifSuffix(puntId puntID, vppSelector string) string {
	return hashString(puntId.String()+":"+vppSelector, 5)
}

// getIcIfaceName returns the name for the interconnect interface (both sides use the same name).
func (m *interconnectManager) getIcIfaceName(ic icID, icType pb.PuntRequest_InterconnectType) string {
	prefix := strings.ToLower(icType.String()) + "-"
	suffix := hashString(ic.String(), 5)
	return prefix + suffix
}

// GetLinuxVrfName returns the name used for Linux VRF device corresponding to the given VPP VRF.
// Method is "static" in the sense that it can be called anytime, regardless of the internal state of the plugin.
func (m *interconnectManager) GetLinuxVrfName(vrf uint32) string {
	return "vrf" + strconv.Itoa(int(vrf))
}

// getIcIfaceMAC (deterministically) generates HW address for the interconnect interface.
func (m *interconnectManager) getIcIfaceMAC(icIfaceName string, vppSide bool) string {
	hwAddr := make(net.HardwareAddr, 6)
	h := fnv.New32a()
	_, _ = h.Write([]byte(icIfaceName))
	hash := h.Sum32()
	hwAddr[0] = 2
	if vppSide {
		hwAddr[1] = 0xfd
	} else {
		hwAddr[1] = 0xfe
	}
	for i := 0; i < 4; i++ {
		hwAddr[i+2] = byte(hash & 0xff)
		hash >>= 8
	}
	return hwAddr.String()
}

// hashString returns a hash of an arbitrarily long string.
// The hash will have <len> characters (shouldn't be more than 7).
func hashString(str string, len int) string {
	const (
		// 32 letters (5 bits to fit single one)
		letters5b = "abcdefghijklmnopqrstuvwxyzABCDEF"
	)
	h := fnv.New32a()
	h.Write([]byte(str))
	hn := h.Sum32()
	var hash string
	bitMask5b := uint32((1 << 5) - 1)
	for i := 0; i < len; i++ {
		hash = string(letters5b[int(hn&bitMask5b)]) + hash
		hn >>= 5
	}
	return hash
}