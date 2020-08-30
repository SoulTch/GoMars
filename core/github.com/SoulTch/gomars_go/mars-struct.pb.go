// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: mars-struct.proto

package gomars_go

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

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{0}
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{1}
}

type AskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AskRequest) Reset() {
	*x = AskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskRequest) ProtoMessage() {}

func (x *AskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskRequest.ProtoReflect.Descriptor instead.
func (*AskRequest) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{2}
}

type AskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AskResponse) Reset() {
	*x = AskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskResponse) ProtoMessage() {}

func (x *AskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskResponse.ProtoReflect.Descriptor instead.
func (*AskResponse) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{3}
}

type NoticeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoticeResponse) Reset() {
	*x = NoticeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeResponse) ProtoMessage() {}

func (x *NoticeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeResponse.ProtoReflect.Descriptor instead.
func (*NoticeResponse) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{4}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_struct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mars_struct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_mars_struct_proto_rawDescGZIP(), []int{5}
}

var File_mars_struct_proto protoreflect.FileDescriptor

var file_mars_struct_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x72, 0x73, 0x2d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c, 0x0a, 0x0a,
	0x41, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x75, 0x6c, 0x54, 0x63, 0x68, 0x2f,
	0x67, 0x6f, 0x6d, 0x61, 0x72, 0x73, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mars_struct_proto_rawDescOnce sync.Once
	file_mars_struct_proto_rawDescData = file_mars_struct_proto_rawDesc
)

func file_mars_struct_proto_rawDescGZIP() []byte {
	file_mars_struct_proto_rawDescOnce.Do(func() {
		file_mars_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_mars_struct_proto_rawDescData)
	})
	return file_mars_struct_proto_rawDescData
}

var file_mars_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mars_struct_proto_goTypes = []interface{}{
	(*ActionRequest)(nil),  // 0: grpc.ActionRequest
	(*ActionResponse)(nil), // 1: grpc.ActionResponse
	(*AskRequest)(nil),     // 2: grpc.AskRequest
	(*AskResponse)(nil),    // 3: grpc.AskResponse
	(*NoticeResponse)(nil), // 4: grpc.NoticeResponse
	(*GetRequest)(nil),     // 5: grpc.GetRequest
}
var file_mars_struct_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mars_struct_proto_init() }
func file_mars_struct_proto_init() {
	if File_mars_struct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mars_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_mars_struct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_mars_struct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskRequest); i {
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
		file_mars_struct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskResponse); i {
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
		file_mars_struct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeResponse); i {
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
		file_mars_struct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
			RawDescriptor: file_mars_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mars_struct_proto_goTypes,
		DependencyIndexes: file_mars_struct_proto_depIdxs,
		MessageInfos:      file_mars_struct_proto_msgTypes,
	}.Build()
	File_mars_struct_proto = out.File
	file_mars_struct_proto_rawDesc = nil
	file_mars_struct_proto_goTypes = nil
	file_mars_struct_proto_depIdxs = nil
}
