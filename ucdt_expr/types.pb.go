// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: cel/types.proto

package ucdt_expr

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

type SourceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source string            `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *SourceData) Reset() {
	*x = SourceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceData) ProtoMessage() {}

func (x *SourceData) ProtoReflect() protoreflect.Message {
	mi := &file_cel_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceData.ProtoReflect.Descriptor instead.
func (*SourceData) Descriptor() ([]byte, []int) {
	return file_cel_types_proto_rawDescGZIP(), []int{0}
}

func (x *SourceData) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SourceData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type SourceDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataset map[string]*SourceData `protobuf:"bytes,1,rep,name=dataset,proto3" json:"dataset,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SourceDataset) Reset() {
	*x = SourceDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cel_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceDataset) ProtoMessage() {}

func (x *SourceDataset) ProtoReflect() protoreflect.Message {
	mi := &file_cel_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceDataset.ProtoReflect.Descriptor instead.
func (*SourceDataset) Descriptor() ([]byte, []int) {
	return file_cel_types_proto_rawDescGZIP(), []int{1}
}

func (x *SourceDataset) GetDataset() map[string]*SourceData {
	if x != nil {
		return x.Dataset
	}
	return nil
}

var File_cel_types_proto protoreflect.FileDescriptor

var file_cel_types_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x63, 0x64, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x63, 0x64, 0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x37,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x99, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x63, 0x64,
	0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x1a, 0x4c, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x63, 0x64, 0x74, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x75, 0x63, 0x64, 0x74, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cel_types_proto_rawDescOnce sync.Once
	file_cel_types_proto_rawDescData = file_cel_types_proto_rawDesc
)

func file_cel_types_proto_rawDescGZIP() []byte {
	file_cel_types_proto_rawDescOnce.Do(func() {
		file_cel_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_cel_types_proto_rawDescData)
	})
	return file_cel_types_proto_rawDescData
}

var file_cel_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cel_types_proto_goTypes = []interface{}{
	(*SourceData)(nil),    // 0: ucdt.SourceData
	(*SourceDataset)(nil), // 1: ucdt.SourceDataset
	nil,                   // 2: ucdt.SourceData.DataEntry
	nil,                   // 3: ucdt.SourceDataset.DatasetEntry
}
var file_cel_types_proto_depIdxs = []int32{
	2, // 0: ucdt.SourceData.data:type_name -> ucdt.SourceData.DataEntry
	3, // 1: ucdt.SourceDataset.dataset:type_name -> ucdt.SourceDataset.DatasetEntry
	0, // 2: ucdt.SourceDataset.DatasetEntry.value:type_name -> ucdt.SourceData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cel_types_proto_init() }
func file_cel_types_proto_init() {
	if File_cel_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cel_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceData); i {
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
		file_cel_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceDataset); i {
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
			RawDescriptor: file_cel_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cel_types_proto_goTypes,
		DependencyIndexes: file_cel_types_proto_depIdxs,
		MessageInfos:      file_cel_types_proto_msgTypes,
	}.Build()
	File_cel_types_proto = out.File
	file_cel_types_proto_rawDesc = nil
	file_cel_types_proto_goTypes = nil
	file_cel_types_proto_depIdxs = nil
}