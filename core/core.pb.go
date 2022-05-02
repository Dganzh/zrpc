// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: core/core.proto

package core

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

type CodecType int32

const (
	CodecType_GOB     CodecType = 0
	CodecType_JSON    CodecType = 1
	CodecType_MSGPACK CodecType = 2
)

// Enum value maps for CodecType.
var (
	CodecType_name = map[int32]string{
		0: "GOB",
		1: "JSON",
		2: "MSGPACK",
	}
	CodecType_value = map[string]int32{
		"GOB":     0,
		"JSON":    1,
		"MSGPACK": 2,
	}
)

func (x CodecType) Enum() *CodecType {
	p := new(CodecType)
	*p = x
	return p
}

func (x CodecType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodecType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_core_proto_enumTypes[0].Descriptor()
}

func (CodecType) Type() protoreflect.EnumType {
	return &file_core_core_proto_enumTypes[0]
}

func (x CodecType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodecType.Descriptor instead.
func (CodecType) EnumDescriptor() ([]byte, []int) {
	return file_core_core_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handler   string    `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	CodecType CodecType `protobuf:"varint,2,opt,name=codecType,proto3,enum=CodecType" json:"codecType,omitempty"`
	Data      []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_core_core_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *Request) GetCodecType() CodecType {
	if x != nil {
		return x.CodecType
	}
	return CodecType_GOB
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The response message containing the greetings
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_core_core_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_core_core_proto protoreflect.FileDescriptor

var file_core_core_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x61, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x2b, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x4f, 0x42, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x53, 0x47, 0x50, 0x41, 0x43, 0x4b, 0x10,
	0x02, 0x32, 0x24, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_core_proto_rawDescOnce sync.Once
	file_core_core_proto_rawDescData = file_core_core_proto_rawDesc
)

func file_core_core_proto_rawDescGZIP() []byte {
	file_core_core_proto_rawDescOnce.Do(func() {
		file_core_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_core_proto_rawDescData)
	})
	return file_core_core_proto_rawDescData
}

var file_core_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_core_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_core_proto_goTypes = []interface{}{
	(CodecType)(0),   // 0: CodecType
	(*Request)(nil),  // 1: Request
	(*Response)(nil), // 2: Response
}
var file_core_core_proto_depIdxs = []int32{
	0, // 0: Request.codecType:type_name -> CodecType
	1, // 1: RPC.Call:input_type -> Request
	2, // 2: RPC.Call:output_type -> Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_core_proto_init() }
func file_core_core_proto_init() {
	if File_core_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_core_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_core_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_core_proto_goTypes,
		DependencyIndexes: file_core_core_proto_depIdxs,
		EnumInfos:         file_core_core_proto_enumTypes,
		MessageInfos:      file_core_core_proto_msgTypes,
	}.Build()
	File_core_core_proto = out.File
	file_core_core_proto_rawDesc = nil
	file_core_core_proto_goTypes = nil
	file_core_core_proto_depIdxs = nil
}
