// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: proto/p.proto

package p

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MessageA) Reset() {
	*x = MessageA{}
	mi := &file_proto_p_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageA) ProtoMessage() {}

func (x *MessageA) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageA.ProtoReflect.Descriptor instead.
func (*MessageA) Descriptor() ([]byte, []int) {
	return file_proto_p_proto_rawDescGZIP(), []int{0}
}

func (x *MessageA) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MessageB) Reset() {
	*x = MessageB{}
	mi := &file_proto_p_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageB) ProtoMessage() {}

func (x *MessageB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageB.ProtoReflect.Descriptor instead.
func (*MessageB) Descriptor() ([]byte, []int) {
	return file_proto_p_proto_rawDescGZIP(), []int{1}
}

func (x *MessageB) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MyWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyMessage *anypb.Any `protobuf:"bytes,1,opt,name=anyMessage,proto3" json:"anyMessage,omitempty"`
}

func (x *MyWrapper) Reset() {
	*x = MyWrapper{}
	mi := &file_proto_p_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyWrapper) ProtoMessage() {}

func (x *MyWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyWrapper.ProtoReflect.Descriptor instead.
func (*MyWrapper) Descriptor() ([]byte, []int) {
	return file_proto_p_proto_rawDescGZIP(), []int{2}
}

func (x *MyWrapper) GetAnyMessage() *anypb.Any {
	if x != nil {
		return x.AnyMessage
	}
	return nil
}

var File_proto_p_proto protoreflect.FileDescriptor

var file_proto_p_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x41,
	0x0a, 0x09, 0x4d, 0x79, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x61,
	0x6e, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x61, 0x6e, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x05, 0x5a, 0x03, 0x2e, 0x2f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_p_proto_rawDescOnce sync.Once
	file_proto_p_proto_rawDescData = file_proto_p_proto_rawDesc
)

func file_proto_p_proto_rawDescGZIP() []byte {
	file_proto_p_proto_rawDescOnce.Do(func() {
		file_proto_p_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_p_proto_rawDescData)
	})
	return file_proto_p_proto_rawDescData
}

var file_proto_p_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_p_proto_goTypes = []any{
	(*MessageA)(nil),  // 0: MessageA
	(*MessageB)(nil),  // 1: MessageB
	(*MyWrapper)(nil), // 2: MyWrapper
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_proto_p_proto_depIdxs = []int32{
	3, // 0: MyWrapper.anyMessage:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_p_proto_init() }
func file_proto_p_proto_init() {
	if File_proto_p_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_p_proto_goTypes,
		DependencyIndexes: file_proto_p_proto_depIdxs,
		MessageInfos:      file_proto_p_proto_msgTypes,
	}.Build()
	File_proto_p_proto = out.File
	file_proto_p_proto_rawDesc = nil
	file_proto_p_proto_goTypes = nil
	file_proto_p_proto_depIdxs = nil
}