// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: time.proto

package grpctime

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetGreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGreetRequest) Reset() {
	*x = GetGreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetRequest) ProtoMessage() {}

func (x *GetGreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetRequest.ProtoReflect.Descriptor instead.
func (*GetGreetRequest) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{1}
}

type GetCurrentTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentTimeRequest) Reset() {
	*x = GetCurrentTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentTimeRequest) ProtoMessage() {}

func (x *GetCurrentTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentTimeRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentTimeRequest) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{2}
}

type GetCurrentTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentTime string `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
}

func (x *GetCurrentTimeResponse) Reset() {
	*x = GetCurrentTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentTimeResponse) ProtoMessage() {}

func (x *GetCurrentTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentTimeResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentTimeResponse) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentTimeResponse) GetCurrentTime() string {
	if x != nil {
		return x.CurrentTime
	}
	return ""
}

type GetGreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
}

func (x *GetGreetResponse) Reset() {
	*x = GetGreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetResponse) ProtoMessage() {}

func (x *GetGreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetResponse.ProtoReflect.Descriptor instead.
func (*GetGreetResponse) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{4}
}

func (x *GetGreetResponse) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_time_proto protoreflect.FileDescriptor

var file_time_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d,
	0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x23,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x32, 0xfd, 0x02, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x70, 0x6c, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x70, 0x6c, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6d, 0x70,
	0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a,
	0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73,
	0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0c, 0x42, 0x61,
	0x63, 0x6b, 0x41, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x19, 0x2e, 0x73, 0x6d, 0x70,
	0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_proto_rawDescOnce sync.Once
	file_time_proto_rawDescData = file_time_proto_rawDesc
)

func file_time_proto_rawDescGZIP() []byte {
	file_time_proto_rawDescOnce.Do(func() {
		file_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_proto_rawDescData)
	})
	return file_time_proto_rawDescData
}

var file_time_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_time_proto_goTypes = []interface{}{
	(*Message)(nil),                // 0: smpl.time.api.v1.Message
	(*GetGreetRequest)(nil),        // 1: smpl.time.api.v1.GetGreetRequest
	(*GetCurrentTimeRequest)(nil),  // 2: smpl.time.api.v1.GetCurrentTimeRequest
	(*GetCurrentTimeResponse)(nil), // 3: smpl.time.api.v1.GetCurrentTimeResponse
	(*GetGreetResponse)(nil),       // 4: smpl.time.api.v1.GetGreetResponse
}
var file_time_proto_depIdxs = []int32{
	2, // 0: smpl.time.api.v1.TimeService.GetCurrentTime:input_type -> smpl.time.api.v1.GetCurrentTimeRequest
	1, // 1: smpl.time.api.v1.TimeService.GetGreet:input_type -> smpl.time.api.v1.GetGreetRequest
	2, // 2: smpl.time.api.v1.TimeService.StreamTimeUpdates:input_type -> smpl.time.api.v1.GetCurrentTimeRequest
	0, // 3: smpl.time.api.v1.TimeService.BackAndForth:input_type -> smpl.time.api.v1.Message
	3, // 4: smpl.time.api.v1.TimeService.GetCurrentTime:output_type -> smpl.time.api.v1.GetCurrentTimeResponse
	4, // 5: smpl.time.api.v1.TimeService.GetGreet:output_type -> smpl.time.api.v1.GetGreetResponse
	3, // 6: smpl.time.api.v1.TimeService.StreamTimeUpdates:output_type -> smpl.time.api.v1.GetCurrentTimeResponse
	0, // 7: smpl.time.api.v1.TimeService.BackAndForth:output_type -> smpl.time.api.v1.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_time_proto_init() }
func file_time_proto_init() {
	if File_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetRequest); i {
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
		file_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentTimeRequest); i {
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
		file_time_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentTimeResponse); i {
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
		file_time_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetResponse); i {
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
			RawDescriptor: file_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_proto_goTypes,
		DependencyIndexes: file_time_proto_depIdxs,
		MessageInfos:      file_time_proto_msgTypes,
	}.Build()
	File_time_proto = out.File
	file_time_proto_rawDesc = nil
	file_time_proto_goTypes = nil
	file_time_proto_depIdxs = nil
}
