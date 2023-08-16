// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.4
// source: upscale.proto

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

type VariationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Index     int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	TaskId    string `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Webhook   string `protobuf:"bytes,4,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *VariationRequest) Reset() {
	*x = VariationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_variation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariationRequest) ProtoMessage() {}

func (x *VariationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_variation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariationRequest.ProtoReflect.Descriptor instead.
func (*VariationRequest) Descriptor() ([]byte, []int) {
	return file_variation_proto_rawDescGZIP(), []int{0}
}

func (x *VariationRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *VariationRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *VariationRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *VariationRequest) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

type VariationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      Codes                  `protobuf:"varint,2,opt,name=code,proto3,enum=api.Codes" json:"code,omitempty"`
	Msg       string                 `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Data      *VariationResponseData `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *VariationResponse) Reset() {
	*x = VariationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_variation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariationResponse) ProtoMessage() {}

func (x *VariationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_variation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariationResponse.ProtoReflect.Descriptor instead.
func (*VariationResponse) Descriptor() ([]byte, []int) {
	return file_variation_proto_rawDescGZIP(), []int{1}
}

func (x *VariationResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *VariationResponse) GetCode() Codes {
	if x != nil {
		return x.Code
	}
	return Codes_CODES_SUCCESS
}

func (x *VariationResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *VariationResponse) GetData() *VariationResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type VariationResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *VariationResponseData) Reset() {
	*x = VariationResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_variation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariationResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariationResponseData) ProtoMessage() {}

func (x *VariationResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_variation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariationResponseData.ProtoReflect.Descriptor instead.
func (*VariationResponseData) Descriptor() ([]byte, []int) {
	return file_variation_proto_rawDescGZIP(), []int{2}
}

func (x *VariationResponseData) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *VariationResponseData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

var File_variation_proto protoreflect.FileDescriptor

var file_variation_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x10, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x22, 0xa2, 0x01, 0x0a, 0x11, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x15, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_variation_proto_rawDescOnce sync.Once
	file_variation_proto_rawDescData = file_variation_proto_rawDesc
)

func file_variation_proto_rawDescGZIP() []byte {
	file_variation_proto_rawDescOnce.Do(func() {
		file_variation_proto_rawDescData = protoimpl.X.CompressGZIP(file_variation_proto_rawDescData)
	})
	return file_variation_proto_rawDescData
}

var file_variation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_variation_proto_goTypes = []interface{}{
	(*VariationRequest)(nil),      // 0: api.VariationRequest
	(*VariationResponse)(nil),     // 1: api.VariationResponse
	(*VariationResponseData)(nil), // 2: api.VariationResponseData
	(Codes)(0),                // 3: api.Codes
}
var file_variation_proto_depIdxs = []int32{
	3, // 0: api.VariationResponse.code:type_name -> api.Codes
	2, // 1: api.VariationResponse.data:type_name -> api.VariationResponseData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_variation_proto_init() }
func file_variation_proto_init() {
	if File_variation_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_variation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariationRequest); i {
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
		file_variation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariationResponse); i {
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
		file_variation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariationResponseData); i {
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
	file_variation_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_variation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_variation_proto_goTypes,
		DependencyIndexes: file_variation_proto_depIdxs,
		MessageInfos:      file_variation_proto_msgTypes,
	}.Build()
	File_variation_proto = out.File
	file_variation_proto_rawDesc = nil
	file_variation_proto_goTypes = nil
	file_variation_proto_depIdxs = nil
}
