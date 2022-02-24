// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: serializers/protoserializer/protostruct.proto

package protoserializer

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

type ProtoSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words   string           `protobuf:"bytes,1,opt,name=words,proto3" json:"words,omitempty"`
	List    []int32          `protobuf:"varint,2,rep,packed,name=list,proto3" json:"list,omitempty"`
	Dict    map[string]int32 `protobuf:"bytes,3,rep,name=dict,proto3" json:"dict,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Integer int32            `protobuf:"varint,4,opt,name=integer,proto3" json:"integer,omitempty"`
	Float   float32          `protobuf:"fixed32,5,opt,name=float,proto3" json:"float,omitempty"`
}

func (x *ProtoSample) Reset() {
	*x = ProtoSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serializers_protoserializer_protostruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoSample) ProtoMessage() {}

func (x *ProtoSample) ProtoReflect() protoreflect.Message {
	mi := &file_serializers_protoserializer_protostruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoSample.ProtoReflect.Descriptor instead.
func (*ProtoSample) Descriptor() ([]byte, []int) {
	return file_serializers_protoserializer_protostruct_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoSample) GetWords() string {
	if x != nil {
		return x.Words
	}
	return ""
}

func (x *ProtoSample) GetList() []int32 {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ProtoSample) GetDict() map[string]int32 {
	if x != nil {
		return x.Dict
	}
	return nil
}

func (x *ProtoSample) GetInteger() int32 {
	if x != nil {
		return x.Integer
	}
	return 0
}

func (x *ProtoSample) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

var File_serializers_protoserializer_protostruct_proto protoreflect.FileDescriptor

var file_serializers_protoserializer_protostruct_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x69, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x64, 0x69, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x69, 0x63, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serializers_protoserializer_protostruct_proto_rawDescOnce sync.Once
	file_serializers_protoserializer_protostruct_proto_rawDescData = file_serializers_protoserializer_protostruct_proto_rawDesc
)

func file_serializers_protoserializer_protostruct_proto_rawDescGZIP() []byte {
	file_serializers_protoserializer_protostruct_proto_rawDescOnce.Do(func() {
		file_serializers_protoserializer_protostruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_serializers_protoserializer_protostruct_proto_rawDescData)
	})
	return file_serializers_protoserializer_protostruct_proto_rawDescData
}

var file_serializers_protoserializer_protostruct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serializers_protoserializer_protostruct_proto_goTypes = []interface{}{
	(*ProtoSample)(nil), // 0: ProtoSample
	nil,                 // 1: ProtoSample.DictEntry
}
var file_serializers_protoserializer_protostruct_proto_depIdxs = []int32{
	1, // 0: ProtoSample.dict:type_name -> ProtoSample.DictEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_serializers_protoserializer_protostruct_proto_init() }
func file_serializers_protoserializer_protostruct_proto_init() {
	if File_serializers_protoserializer_protostruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serializers_protoserializer_protostruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoSample); i {
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
			RawDescriptor: file_serializers_protoserializer_protostruct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serializers_protoserializer_protostruct_proto_goTypes,
		DependencyIndexes: file_serializers_protoserializer_protostruct_proto_depIdxs,
		MessageInfos:      file_serializers_protoserializer_protostruct_proto_msgTypes,
	}.Build()
	File_serializers_protoserializer_protostruct_proto = out.File
	file_serializers_protoserializer_protostruct_proto_rawDesc = nil
	file_serializers_protoserializer_protostruct_proto_goTypes = nil
	file_serializers_protoserializer_protostruct_proto_depIdxs = nil
}
