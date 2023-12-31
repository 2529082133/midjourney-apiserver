// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.0
// source: common.proto

package api

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

type Codes int32

const (
	Codes_CODES_SUCCESS                 Codes = 0
	Codes_CODES_SERVER_INTERNAL_ERROR   Codes = 1
	Codes_CODES_INVALID_PARAMETER_ERROR Codes = 2
	Codes_CODES_CONCURRENCY_LIMITED     Codes = 3
	Codes_CODES_PROCESSING_TIMEOUT      Codes = 4
)

// Enum value maps for Codes.
var (
	Codes_name = map[int32]string{
		0: "CODES_SUCCESS",
		1: "CODES_SERVER_INTERNAL_ERROR",
		2: "CODES_INVALID_PARAMETER_ERROR",
		3: "CODES_CONCURRENCY_LIMITED",
		4: "CODES_PROCESSING_TIMEOUT",
	}
	Codes_value = map[string]int32{
		"CODES_SUCCESS":                 0,
		"CODES_SERVER_INTERNAL_ERROR":   1,
		"CODES_INVALID_PARAMETER_ERROR": 2,
		"CODES_CONCURRENCY_LIMITED":     3,
		"CODES_PROCESSING_TIMEOUT":      4,
	}
)

func (x Codes) Enum() *Codes {
	p := new(Codes)
	*p = x
	return p
}

func (x Codes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Codes) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Codes) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Codes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Codes.Descriptor instead.
func (Codes) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x2a, 0x9b, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x43, 0x4f,
	0x4e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x04, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_goTypes = []interface{}{
	(Codes)(0), // 0: api.Codes
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
