// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: table_records.proto

package mysql

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

type PbField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	KeyType int32  `protobuf:"varint,2,opt,name=KeyType,proto3" json:"KeyType,omitempty"`
	Type    int32  `protobuf:"zigzag32,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Value   []byte `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *PbField) Reset() {
	*x = PbField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_records_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbField) ProtoMessage() {}

func (x *PbField) ProtoReflect() protoreflect.Message {
	mi := &file_table_records_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbField.ProtoReflect.Descriptor instead.
func (*PbField) Descriptor() ([]byte, []int) {
	return file_table_records_proto_rawDescGZIP(), []int{0}
}

func (x *PbField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PbField) GetKeyType() int32 {
	if x != nil {
		return x.KeyType
	}
	return 0
}

func (x *PbField) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PbField) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PbRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*PbField `protobuf:"bytes,1,rep,name=Fields,proto3" json:"Fields,omitempty"`
}

func (x *PbRow) Reset() {
	*x = PbRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_records_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbRow) ProtoMessage() {}

func (x *PbRow) ProtoReflect() protoreflect.Message {
	mi := &file_table_records_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbRow.ProtoReflect.Descriptor instead.
func (*PbRow) Descriptor() ([]byte, []int) {
	return file_table_records_proto_rawDescGZIP(), []int{1}
}

func (x *PbRow) GetFields() []*PbField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type PbTableRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string   `protobuf:"bytes,1,opt,name=TableName,proto3" json:"TableName,omitempty"`
	Rows      []*PbRow `protobuf:"bytes,2,rep,name=Rows,proto3" json:"Rows,omitempty"`
}

func (x *PbTableRecords) Reset() {
	*x = PbTableRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_records_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbTableRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbTableRecords) ProtoMessage() {}

func (x *PbTableRecords) ProtoReflect() protoreflect.Message {
	mi := &file_table_records_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbTableRecords.ProtoReflect.Descriptor instead.
func (*PbTableRecords) Descriptor() ([]byte, []int) {
	return file_table_records_proto_rawDescGZIP(), []int{2}
}

func (x *PbTableRecords) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *PbTableRecords) GetRows() []*PbRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_table_records_proto protoreflect.FileDescriptor

var file_table_records_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x22, 0x61, 0x0a, 0x07,
	0x50, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4b, 0x65,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2f, 0x0a, 0x05, 0x50, 0x62, 0x52, 0x6f, 0x77, 0x12, 0x26, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x2e, 0x50, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x22, 0x50, 0x0a, 0x0e, 0x50, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x50, 0x62, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x52, 0x6f,
	0x77, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_table_records_proto_rawDescOnce sync.Once
	file_table_records_proto_rawDescData = file_table_records_proto_rawDesc
)

func file_table_records_proto_rawDescGZIP() []byte {
	file_table_records_proto_rawDescOnce.Do(func() {
		file_table_records_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_records_proto_rawDescData)
	})
	return file_table_records_proto_rawDescData
}

var file_table_records_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_table_records_proto_goTypes = []interface{}{
	(*PbField)(nil),        // 0: mysql.PbField
	(*PbRow)(nil),          // 1: mysql.PbRow
	(*PbTableRecords)(nil), // 2: mysql.PbTableRecords
}
var file_table_records_proto_depIdxs = []int32{
	0, // 0: mysql.PbRow.Fields:type_name -> mysql.PbField
	1, // 1: mysql.PbTableRecords.Rows:type_name -> mysql.PbRow
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_table_records_proto_init() }
func file_table_records_proto_init() {
	if File_table_records_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_records_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbField); i {
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
		file_table_records_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbRow); i {
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
		file_table_records_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbTableRecords); i {
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
			RawDescriptor: file_table_records_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_table_records_proto_goTypes,
		DependencyIndexes: file_table_records_proto_depIdxs,
		MessageInfos:      file_table_records_proto_msgTypes,
	}.Build()
	File_table_records_proto = out.File
	file_table_records_proto_rawDesc = nil
	file_table_records_proto_goTypes = nil
	file_table_records_proto_depIdxs = nil
}