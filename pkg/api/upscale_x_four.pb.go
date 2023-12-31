// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.0
// source: upscale_x_four.proto

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

type UpscaleXFourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Index     int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	TaskId    string `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Webhook   string `protobuf:"bytes,4,opt,name=webhook,proto3" json:"webhook,omitempty"`
	MemberId  string `protobuf:"bytes,5,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *UpscaleXFourRequest) Reset() {
	*x = UpscaleXFourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upscale_x_four_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleXFourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleXFourRequest) ProtoMessage() {}

func (x *UpscaleXFourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upscale_x_four_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleXFourRequest.ProtoReflect.Descriptor instead.
func (*UpscaleXFourRequest) Descriptor() ([]byte, []int) {
	return file_upscale_x_four_proto_rawDescGZIP(), []int{0}
}

func (x *UpscaleXFourRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpscaleXFourRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UpscaleXFourRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpscaleXFourRequest) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

func (x *UpscaleXFourRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type UpscaleXFourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                    `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      Codes                     `protobuf:"varint,2,opt,name=code,proto3,enum=api.Codes" json:"code,omitempty"`
	Msg       string                    `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Data      *UpscaleXFourResponseData `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *UpscaleXFourResponse) Reset() {
	*x = UpscaleXFourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upscale_x_four_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleXFourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleXFourResponse) ProtoMessage() {}

func (x *UpscaleXFourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upscale_x_four_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleXFourResponse.ProtoReflect.Descriptor instead.
func (*UpscaleXFourResponse) Descriptor() ([]byte, []int) {
	return file_upscale_x_four_proto_rawDescGZIP(), []int{1}
}

func (x *UpscaleXFourResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpscaleXFourResponse) GetCode() Codes {
	if x != nil {
		return x.Code
	}
	return Codes_CODES_SUCCESS
}

func (x *UpscaleXFourResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpscaleXFourResponse) GetData() *UpscaleXFourResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpscaleXFourResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *UpscaleXFourResponseData) Reset() {
	*x = UpscaleXFourResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upscale_x_four_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleXFourResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleXFourResponseData) ProtoMessage() {}

func (x *UpscaleXFourResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_upscale_x_four_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleXFourResponseData.ProtoReflect.Descriptor instead.
func (*UpscaleXFourResponseData) Descriptor() ([]byte, []int) {
	return file_upscale_x_four_proto_rawDescGZIP(), []int{2}
}

func (x *UpscaleXFourResponseData) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpscaleXFourResponseData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

var File_upscale_x_four_proto protoreflect.FileDescriptor

var file_upscale_x_four_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x78, 0x5f, 0x66, 0x6f, 0x75, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x55, 0x70,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x58, 0x46, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x58, 0x46, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x58, 0x46, 0x6f, 0x75,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x52, 0x0a, 0x18, 0x55, 0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x58, 0x46, 0x6f, 0x75,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upscale_x_four_proto_rawDescOnce sync.Once
	file_upscale_x_four_proto_rawDescData = file_upscale_x_four_proto_rawDesc
)

func file_upscale_x_four_proto_rawDescGZIP() []byte {
	file_upscale_x_four_proto_rawDescOnce.Do(func() {
		file_upscale_x_four_proto_rawDescData = protoimpl.X.CompressGZIP(file_upscale_x_four_proto_rawDescData)
	})
	return file_upscale_x_four_proto_rawDescData
}

var file_upscale_x_four_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_upscale_x_four_proto_goTypes = []interface{}{
	(*UpscaleXFourRequest)(nil),      // 0: api.UpscaleXFourRequest
	(*UpscaleXFourResponse)(nil),     // 1: api.UpscaleXFourResponse
	(*UpscaleXFourResponseData)(nil), // 2: api.UpscaleXFourResponseData
	(Codes)(0),                       // 3: api.Codes
}
var file_upscale_x_four_proto_depIdxs = []int32{
	3, // 0: api.UpscaleXFourResponse.code:type_name -> api.Codes
	2, // 1: api.UpscaleXFourResponse.data:type_name -> api.UpscaleXFourResponseData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upscale_x_four_proto_init() }
func file_upscale_x_four_proto_init() {
	if File_upscale_x_four_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_upscale_x_four_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleXFourRequest); i {
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
		file_upscale_x_four_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleXFourResponse); i {
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
		file_upscale_x_four_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleXFourResponseData); i {
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
	file_upscale_x_four_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upscale_x_four_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upscale_x_four_proto_goTypes,
		DependencyIndexes: file_upscale_x_four_proto_depIdxs,
		MessageInfos:      file_upscale_x_four_proto_msgTypes,
	}.Build()
	File_upscale_x_four_proto = out.File
	file_upscale_x_four_proto_rawDesc = nil
	file_upscale_x_four_proto_goTypes = nil
	file_upscale_x_four_proto_depIdxs = nil
}
