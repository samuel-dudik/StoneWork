// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cnfreg/cnfreg.proto

package cnfreg

import (
	proto "github.com/golang/protobuf/proto"
	generic "go.ligato.io/vpp-agent/v3/proto/ligato/generic"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	puntmgr "go.pantheon.tech/stonework/proto/puntmgr"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Mode in which the CNF operates wrt. the other CNFs.
type CnfMode int32

const (
	// CNF runs both data-plane (typically VPP) and control-plane inside its container
	// and with other CNFs is integrated through CNF chaining (i.e. packets are copied between CNFs).
	CnfMode_STANDALONE CnfMode = 0
	// CNF runs as a module for StoneWork (All-in-one VPP). It means that it does not run data-plane inside
	// its container, but instead is integrated with the single shared VPP and potentially also shares
	// the same Linux network namespace with some other CNFs.
	CnfMode_STONEWORK_MODULE CnfMode = 1
	// Mode in which the agent of StoneWork operates.
	CnfMode_STONEWORK CnfMode = 2
)

// Enum value maps for CnfMode.
var (
	CnfMode_name = map[int32]string{
		0: "STANDALONE",
		1: "STONEWORK_MODULE",
		2: "STONEWORK",
	}
	CnfMode_value = map[string]int32{
		"STANDALONE":       0,
		"STONEWORK_MODULE": 1,
		"STONEWORK":        2,
	}
)

func (x CnfMode) Enum() *CnfMode {
	p := new(CnfMode)
	*p = x
	return p
}

func (x CnfMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CnfMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cnfreg_cnfreg_proto_enumTypes[0].Descriptor()
}

func (CnfMode) Type() protoreflect.EnumType {
	return &file_cnfreg_cnfreg_proto_enumTypes[0]
}

func (x CnfMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CnfMode.Descriptor instead.
func (CnfMode) EnumDescriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{0}
}

// DiscoverCnfReq is sent by CNFRegistry of STONEWORK to discover a SW-Module CNF.
type DiscoverCnfReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Management IP address of StoneWork.
	SwIpAddress string `protobuf:"bytes,1,opt,name=sw_ip_address,json=swIpAddress,proto3" json:"sw_ip_address,omitempty"`
	// gRPC port on which StoneWork (client of this request) listens.
	SwGrpcPort uint32 `protobuf:"varint,2,opt,name=sw_grpc_port,json=swGrpcPort,proto3" json:"sw_grpc_port,omitempty"`
	// HTTP port on which StoneWork (client of this request) listens.
	SwHttpPort uint32 `protobuf:"varint,3,opt,name=sw_http_port,json=swHttpPort,proto3" json:"sw_http_port,omitempty"`
}

func (x *DiscoverCnfReq) Reset() {
	*x = DiscoverCnfReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverCnfReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverCnfReq) ProtoMessage() {}

func (x *DiscoverCnfReq) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverCnfReq.ProtoReflect.Descriptor instead.
func (*DiscoverCnfReq) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoverCnfReq) GetSwIpAddress() string {
	if x != nil {
		return x.SwIpAddress
	}
	return ""
}

func (x *DiscoverCnfReq) GetSwGrpcPort() uint32 {
	if x != nil {
		return x.SwGrpcPort
	}
	return 0
}

func (x *DiscoverCnfReq) GetSwHttpPort() uint32 {
	if x != nil {
		return x.SwHttpPort
	}
	return 0
}

// DiscoverCnfResp is returned by STONEWORK_MODULE with information about CNF configuration models.
type DiscoverCnfResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Microservice label of the discovered CNF.
	CnfMsLabel   string                         `protobuf:"bytes,1,opt,name=cnf_ms_label,json=cnfMsLabel,proto3" json:"cnf_ms_label,omitempty"`
	ConfigModels []*DiscoverCnfResp_ConfigModel `protobuf:"bytes,4,rep,name=config_models,json=configModels,proto3" json:"config_models,omitempty"`
}

func (x *DiscoverCnfResp) Reset() {
	*x = DiscoverCnfResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverCnfResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverCnfResp) ProtoMessage() {}

func (x *DiscoverCnfResp) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverCnfResp.ProtoReflect.Descriptor instead.
func (*DiscoverCnfResp) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoverCnfResp) GetCnfMsLabel() string {
	if x != nil {
		return x.CnfMsLabel
	}
	return ""
}

func (x *DiscoverCnfResp) GetConfigModels() []*DiscoverCnfResp_ConfigModel {
	if x != nil {
		return x.ConfigModels
	}
	return nil
}

// ConfigItemDependency stores information about a single dependency of a configuration item.
type ConfigItemDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Types that are assignable to Dep:
	//	*ConfigItemDependency_Key_
	//	*ConfigItemDependency_Anyof
	Dep isConfigItemDependency_Dep `protobuf_oneof:"dep"`
}

func (x *ConfigItemDependency) Reset() {
	*x = ConfigItemDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigItemDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItemDependency) ProtoMessage() {}

func (x *ConfigItemDependency) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItemDependency.ProtoReflect.Descriptor instead.
func (*ConfigItemDependency) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigItemDependency) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (m *ConfigItemDependency) GetDep() isConfigItemDependency_Dep {
	if m != nil {
		return m.Dep
	}
	return nil
}

func (x *ConfigItemDependency) GetKey() string {
	if x, ok := x.GetDep().(*ConfigItemDependency_Key_); ok {
		return x.Key
	}
	return ""
}

func (x *ConfigItemDependency) GetAnyof() *ConfigItemDependency_AnyOf {
	if x, ok := x.GetDep().(*ConfigItemDependency_Anyof); ok {
		return x.Anyof
	}
	return nil
}

type isConfigItemDependency_Dep interface {
	isConfigItemDependency_Dep()
}

type ConfigItemDependency_Key_ struct {
	Key string `protobuf:"bytes,2,opt,name=key,proto3,oneof"`
}

type ConfigItemDependency_Anyof struct {
	Anyof *ConfigItemDependency_AnyOf `protobuf:"bytes,3,opt,name=anyof,proto3,oneof"`
}

func (*ConfigItemDependency_Key_) isConfigItemDependency_Dep() {}

func (*ConfigItemDependency_Anyof) isConfigItemDependency_Dep() {}

// GetDependenciesResp is returned by STONEWORK_MODULE to inform about dependencies
// of a given configuration item.
type GetDependenciesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dependencies []*ConfigItemDependency `protobuf:"bytes,1,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *GetDependenciesResp) Reset() {
	*x = GetDependenciesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDependenciesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDependenciesResp) ProtoMessage() {}

func (x *GetDependenciesResp) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDependenciesResp.ProtoReflect.Descriptor instead.
func (*GetDependenciesResp) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{3}
}

func (x *GetDependenciesResp) GetDependencies() []*ConfigItemDependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type DiscoverCnfResp_ConfigModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ProtoName is a name of protobuf message representing the model.
	ProtoName    string `protobuf:"bytes,1,opt,name=proto_name,json=protoName,proto3" json:"proto_name,omitempty"`
	WithPunt     bool   `protobuf:"varint,2,opt,name=with_punt,json=withPunt,proto3" json:"with_punt,omitempty"`
	WithRetrieve bool   `protobuf:"varint,3,opt,name=with_retrieve,json=withRetrieve,proto3" json:"with_retrieve,omitempty"`
	WithDeps     bool   `protobuf:"varint,4,opt,name=with_deps,json=withDeps,proto3" json:"with_deps,omitempty"`
}

func (x *DiscoverCnfResp_ConfigModel) Reset() {
	*x = DiscoverCnfResp_ConfigModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverCnfResp_ConfigModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverCnfResp_ConfigModel) ProtoMessage() {}

func (x *DiscoverCnfResp_ConfigModel) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverCnfResp_ConfigModel.ProtoReflect.Descriptor instead.
func (*DiscoverCnfResp_ConfigModel) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DiscoverCnfResp_ConfigModel) GetProtoName() string {
	if x != nil {
		return x.ProtoName
	}
	return ""
}

func (x *DiscoverCnfResp_ConfigModel) GetWithPunt() bool {
	if x != nil {
		return x.WithPunt
	}
	return false
}

func (x *DiscoverCnfResp_ConfigModel) GetWithRetrieve() bool {
	if x != nil {
		return x.WithRetrieve
	}
	return false
}

func (x *DiscoverCnfResp_ConfigModel) GetWithDeps() bool {
	if x != nil {
		return x.WithDeps
	}
	return false
}

type ConfigItemDependency_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ConfigItemDependency_Key) Reset() {
	*x = ConfigItemDependency_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigItemDependency_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItemDependency_Key) ProtoMessage() {}

func (x *ConfigItemDependency_Key) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItemDependency_Key.ProtoReflect.Descriptor instead.
func (*ConfigItemDependency_Key) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ConfigItemDependency_Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ConfigItemDependency_AnyOf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyPrefixes []string `protobuf:"bytes,1,rep,name=key_prefixes,json=keyPrefixes,proto3" json:"key_prefixes,omitempty"`
}

func (x *ConfigItemDependency_AnyOf) Reset() {
	*x = ConfigItemDependency_AnyOf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnfreg_cnfreg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigItemDependency_AnyOf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItemDependency_AnyOf) ProtoMessage() {}

func (x *ConfigItemDependency_AnyOf) ProtoReflect() protoreflect.Message {
	mi := &file_cnfreg_cnfreg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItemDependency_AnyOf.ProtoReflect.Descriptor instead.
func (*ConfigItemDependency_AnyOf) Descriptor() ([]byte, []int) {
	return file_cnfreg_cnfreg_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ConfigItemDependency_AnyOf) GetKeyPrefixes() []string {
	if x != nil {
		return x.KeyPrefixes
	}
	return nil
}

var File_cnfreg_cnfreg_proto protoreflect.FileDescriptor

var file_cnfreg_cnfreg_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x2f, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x1a, 0x15, 0x70,
	0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x70, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6e,
	0x66, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x77, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x77, 0x49,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x77, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x77, 0x47, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x77,
	0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x77, 0x48, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x8b, 0x02, 0x0a,
	0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6e, 0x66, 0x5f, 0x6d, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6e, 0x66, 0x4d, 0x73, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x48, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6e, 0x66, 0x72,
	0x65, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6e, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x8b, 0x01, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x70, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x50, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x77, 0x69, 0x74, 0x68, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x44, 0x65, 0x70, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x14, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a,
	0x05, 0x61, 0x6e, 0x79, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x41, 0x6e, 0x79, 0x4f, 0x66,
	0x48, 0x00, 0x52, 0x05, 0x61, 0x6e, 0x79, 0x6f, 0x66, 0x1a, 0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x1a, 0x2a, 0x0a, 0x05, 0x41, 0x6e, 0x79, 0x4f, 0x66, 0x12, 0x21, 0x0a, 0x0c, 0x6b,
	0x65, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x42, 0x05,
	0x0a, 0x03, 0x64, 0x65, 0x70, 0x22, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0c,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2a, 0x3e,
	0x0a, 0x07, 0x43, 0x6e, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41,
	0x4e, 0x44, 0x41, 0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x4f,
	0x4e, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x54, 0x4f, 0x4e, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x32, 0xd8,
	0x01, 0x0a, 0x0c, 0x43, 0x6e, 0x66, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x3e, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6e, 0x66, 0x12, 0x16,
	0x2e, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x43, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15, 0x2e, 0x70, 0x75, 0x6e, 0x74, 0x6d,
	0x67, 0x72, 0x2e, 0x50, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x1b, 0x2e, 0x63,
	0x6e, 0x66, 0x72, 0x65, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x2d, 0x5a, 0x2b, 0x70, 0x61, 0x6e,
	0x74, 0x68, 0x65, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6e, 0x66, 0x72, 0x65,
	0x67, 0x3b, 0x63, 0x6e, 0x66, 0x72, 0x65, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cnfreg_cnfreg_proto_rawDescOnce sync.Once
	file_cnfreg_cnfreg_proto_rawDescData = file_cnfreg_cnfreg_proto_rawDesc
)

func file_cnfreg_cnfreg_proto_rawDescGZIP() []byte {
	file_cnfreg_cnfreg_proto_rawDescOnce.Do(func() {
		file_cnfreg_cnfreg_proto_rawDescData = protoimpl.X.CompressGZIP(file_cnfreg_cnfreg_proto_rawDescData)
	})
	return file_cnfreg_cnfreg_proto_rawDescData
}

var file_cnfreg_cnfreg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cnfreg_cnfreg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cnfreg_cnfreg_proto_goTypes = []interface{}{
	(CnfMode)(0),                        // 0: cnfreg.CnfMode
	(*DiscoverCnfReq)(nil),              // 1: cnfreg.DiscoverCnfReq
	(*DiscoverCnfResp)(nil),             // 2: cnfreg.DiscoverCnfResp
	(*ConfigItemDependency)(nil),        // 3: cnfreg.ConfigItemDependency
	(*GetDependenciesResp)(nil),         // 4: cnfreg.GetDependenciesResp
	(*DiscoverCnfResp_ConfigModel)(nil), // 5: cnfreg.DiscoverCnfResp.ConfigModel
	(*ConfigItemDependency_Key)(nil),    // 6: cnfreg.ConfigItemDependency.Key
	(*ConfigItemDependency_AnyOf)(nil),  // 7: cnfreg.ConfigItemDependency.AnyOf
	(*generic.Item)(nil),                // 8: ligato.generic.Item
	(*puntmgr.PuntRequests)(nil),        // 9: puntmgr.PuntRequests
}
var file_cnfreg_cnfreg_proto_depIdxs = []int32{
	5, // 0: cnfreg.DiscoverCnfResp.config_models:type_name -> cnfreg.DiscoverCnfResp.ConfigModel
	7, // 1: cnfreg.ConfigItemDependency.anyof:type_name -> cnfreg.ConfigItemDependency.AnyOf
	3, // 2: cnfreg.GetDependenciesResp.dependencies:type_name -> cnfreg.ConfigItemDependency
	1, // 3: cnfreg.CnfDiscovery.DiscoverCnf:input_type -> cnfreg.DiscoverCnfReq
	8, // 4: cnfreg.CnfDiscovery.GetPuntRequests:input_type -> ligato.generic.Item
	8, // 5: cnfreg.CnfDiscovery.GetItemDependencies:input_type -> ligato.generic.Item
	2, // 6: cnfreg.CnfDiscovery.DiscoverCnf:output_type -> cnfreg.DiscoverCnfResp
	9, // 7: cnfreg.CnfDiscovery.GetPuntRequests:output_type -> puntmgr.PuntRequests
	4, // 8: cnfreg.CnfDiscovery.GetItemDependencies:output_type -> cnfreg.GetDependenciesResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cnfreg_cnfreg_proto_init() }
func file_cnfreg_cnfreg_proto_init() {
	if File_cnfreg_cnfreg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cnfreg_cnfreg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverCnfReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverCnfResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigItemDependency); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDependenciesResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverCnfResp_ConfigModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigItemDependency_Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cnfreg_cnfreg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigItemDependency_AnyOf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cnfreg_cnfreg_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ConfigItemDependency_Key_)(nil),
		(*ConfigItemDependency_Anyof)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cnfreg_cnfreg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cnfreg_cnfreg_proto_goTypes,
		DependencyIndexes: file_cnfreg_cnfreg_proto_depIdxs,
		EnumInfos:         file_cnfreg_cnfreg_proto_enumTypes,
		MessageInfos:      file_cnfreg_cnfreg_proto_msgTypes,
	}.Build()
	File_cnfreg_cnfreg_proto = out.File
	file_cnfreg_cnfreg_proto_rawDesc = nil
	file_cnfreg_cnfreg_proto_goTypes = nil
	file_cnfreg_cnfreg_proto_depIdxs = nil
}
