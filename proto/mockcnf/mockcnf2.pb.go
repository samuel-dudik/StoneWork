// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: mockcnf/mockcnf2.proto

package mockcnf

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// MockCnf is used for testing of CNFRegistry and PuntManager.
// Note: two different CNFs cannot use proto messages from the same proto files,
//       because otherwise two distinct copies of the same file descriptor would be
//       received over separate gRPC calls to KnownModels service, while only
//       one descriptor can be registered for the same proto file.
type MockCnf2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VppInterface string `protobuf:"bytes,1,opt,name=vpp_interface,json=vppInterface,proto3" json:"vpp_interface,omitempty"`
	Vrf          uint32 `protobuf:"varint,2,opt,name=vrf,proto3" json:"vrf,omitempty"`
}

func (x *MockCnf2) Reset() {
	*x = MockCnf2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockcnf_mockcnf2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockCnf2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockCnf2) ProtoMessage() {}

func (x *MockCnf2) ProtoReflect() protoreflect.Message {
	mi := &file_mockcnf_mockcnf2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockCnf2.ProtoReflect.Descriptor instead.
func (*MockCnf2) Descriptor() ([]byte, []int) {
	return file_mockcnf_mockcnf2_proto_rawDescGZIP(), []int{0}
}

func (x *MockCnf2) GetVppInterface() string {
	if x != nil {
		return x.VppInterface
	}
	return ""
}

func (x *MockCnf2) GetVrf() uint32 {
	if x != nil {
		return x.Vrf
	}
	return 0
}

var File_mockcnf_mockcnf2_proto protoreflect.FileDescriptor

var file_mockcnf_mockcnf2_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x63, 0x6b, 0x63, 0x6e, 0x66, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x63, 0x6e,
	0x66, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x6f, 0x63, 0x6b, 0x63, 0x6e,
	0x66, 0x22, 0x41, 0x0a, 0x08, 0x4d, 0x6f, 0x63, 0x6b, 0x43, 0x6e, 0x66, 0x32, 0x12, 0x23, 0x0a,
	0x0d, 0x76, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x72, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x76, 0x72, 0x66, 0x42, 0x2f, 0x5a, 0x2d, 0x70, 0x61, 0x6e, 0x74, 0x68, 0x65, 0x6f, 0x6e,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x63, 0x6e, 0x66, 0x3b, 0x6d, 0x6f,
	0x63, 0x6b, 0x63, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockcnf_mockcnf2_proto_rawDescOnce sync.Once
	file_mockcnf_mockcnf2_proto_rawDescData = file_mockcnf_mockcnf2_proto_rawDesc
)

func file_mockcnf_mockcnf2_proto_rawDescGZIP() []byte {
	file_mockcnf_mockcnf2_proto_rawDescOnce.Do(func() {
		file_mockcnf_mockcnf2_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockcnf_mockcnf2_proto_rawDescData)
	})
	return file_mockcnf_mockcnf2_proto_rawDescData
}

var file_mockcnf_mockcnf2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mockcnf_mockcnf2_proto_goTypes = []interface{}{
	(*MockCnf2)(nil), // 0: mockcnf.MockCnf2
}
var file_mockcnf_mockcnf2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mockcnf_mockcnf2_proto_init() }
func file_mockcnf_mockcnf2_proto_init() {
	if File_mockcnf_mockcnf2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockcnf_mockcnf2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockCnf2); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockcnf_mockcnf2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockcnf_mockcnf2_proto_goTypes,
		DependencyIndexes: file_mockcnf_mockcnf2_proto_depIdxs,
		MessageInfos:      file_mockcnf_mockcnf2_proto_msgTypes,
	}.Build()
	File_mockcnf_mockcnf2_proto = out.File
	file_mockcnf_mockcnf2_proto_rawDesc = nil
	file_mockcnf_mockcnf2_proto_goTypes = nil
	file_mockcnf_mockcnf2_proto_depIdxs = nil
}
