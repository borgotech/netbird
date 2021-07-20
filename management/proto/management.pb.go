// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: management.proto

package proto

import (
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

type RegisterPeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Wireguard public key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Pre-authorized setup key
	SetupKey string `protobuf:"bytes,2,opt,name=setupKey,proto3" json:"setupKey,omitempty"`
}

func (x *RegisterPeerRequest) Reset() {
	*x = RegisterPeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerRequest) ProtoMessage() {}

func (x *RegisterPeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerRequest.ProtoReflect.Descriptor instead.
func (*RegisterPeerRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterPeerRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RegisterPeerRequest) GetSetupKey() string {
	if x != nil {
		return x.SetupKey
	}
	return ""
}

type RegisterPeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterPeerResponse) Reset() {
	*x = RegisterPeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerResponse) ProtoMessage() {}

func (x *RegisterPeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerResponse.ProtoReflect.Descriptor instead.
func (*RegisterPeerResponse) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{2}
}

var File_management_proto protoreflect.FileDescriptor

var file_management_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x43,
	0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x4b, 0x65, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0x9d, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x09, 0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x11, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x11, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_proto_rawDescOnce sync.Once
	file_management_proto_rawDescData = file_management_proto_rawDesc
)

func file_management_proto_rawDescGZIP() []byte {
	file_management_proto_rawDescOnce.Do(func() {
		file_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_proto_rawDescData)
	})
	return file_management_proto_rawDescData
}

var file_management_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_management_proto_goTypes = []interface{}{
	(*RegisterPeerRequest)(nil),  // 0: management.RegisterPeerRequest
	(*RegisterPeerResponse)(nil), // 1: management.RegisterPeerResponse
	(*Empty)(nil),                // 2: management.Empty
}
var file_management_proto_depIdxs = []int32{
	0, // 0: management.ManagementService.RegisterPeer:input_type -> management.RegisterPeerRequest
	2, // 1: management.ManagementService.isHealthy:input_type -> management.Empty
	1, // 2: management.ManagementService.RegisterPeer:output_type -> management.RegisterPeerResponse
	2, // 3: management.ManagementService.isHealthy:output_type -> management.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_proto_init() }
func file_management_proto_init() {
	if File_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerRequest); i {
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
		file_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerResponse); i {
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
		file_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_proto_goTypes,
		DependencyIndexes: file_management_proto_depIdxs,
		MessageInfos:      file_management_proto_msgTypes,
	}.Build()
	File_management_proto = out.File
	file_management_proto_rawDesc = nil
	file_management_proto_goTypes = nil
	file_management_proto_depIdxs = nil
}
