// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: dslog/dslog.proto

package dslog

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

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dslog_dslog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dslog_dslog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_dslog_dslog_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dslog_dslog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dslog_dslog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_dslog_dslog_proto_rawDescGZIP(), []int{1}
}

func (x *LogResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_dslog_dslog_proto protoreflect.FileDescriptor

var file_dslog_dslog_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x64, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x73, 0x6c, 0x6f, 0x67, 0x22, 0x26, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x3a, 0x0a, 0x05, 0x44, 0x73, 0x6c,
	0x6f, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x64,
	0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x64, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x79, 0x77, 0x2f, 0x64, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x64,
	0x73, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dslog_dslog_proto_rawDescOnce sync.Once
	file_dslog_dslog_proto_rawDescData = file_dslog_dslog_proto_rawDesc
)

func file_dslog_dslog_proto_rawDescGZIP() []byte {
	file_dslog_dslog_proto_rawDescOnce.Do(func() {
		file_dslog_dslog_proto_rawDescData = protoimpl.X.CompressGZIP(file_dslog_dslog_proto_rawDescData)
	})
	return file_dslog_dslog_proto_rawDescData
}

var file_dslog_dslog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dslog_dslog_proto_goTypes = []interface{}{
	(*LogRequest)(nil),  // 0: dslog.LogRequest
	(*LogResponse)(nil), // 1: dslog.LogResponse
}
var file_dslog_dslog_proto_depIdxs = []int32{
	0, // 0: dslog.Dslog.AddLog:input_type -> dslog.LogRequest
	1, // 1: dslog.Dslog.AddLog:output_type -> dslog.LogResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dslog_dslog_proto_init() }
func file_dslog_dslog_proto_init() {
	if File_dslog_dslog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dslog_dslog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_dslog_dslog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
			RawDescriptor: file_dslog_dslog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dslog_dslog_proto_goTypes,
		DependencyIndexes: file_dslog_dslog_proto_depIdxs,
		MessageInfos:      file_dslog_dslog_proto_msgTypes,
	}.Build()
	File_dslog_dslog_proto = out.File
	file_dslog_dslog_proto_rawDesc = nil
	file_dslog_dslog_proto_goTypes = nil
	file_dslog_dslog_proto_depIdxs = nil
}