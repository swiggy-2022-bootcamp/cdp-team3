// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: send.proto

package proto

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

type RewardDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Reward int32  `protobuf:"varint,2,opt,name=Reward,proto3" json:"Reward,omitempty"`
}

func (x *RewardDetails) Reset() {
	*x = RewardDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardDetails) ProtoMessage() {}

func (x *RewardDetails) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardDetails.ProtoReflect.Descriptor instead.
func (*RewardDetails) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{0}
}

func (x *RewardDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RewardDetails) GetReward() int32 {
	if x != nil {
		return x.Reward
	}
	return 0
}

type SuccessMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdded string `protobuf:"bytes,1,opt,name=IsAdded,proto3" json:"IsAdded,omitempty"`
}

func (x *SuccessMessage) Reset() {
	*x = SuccessMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessMessage) ProtoMessage() {}

func (x *SuccessMessage) ProtoReflect() protoreflect.Message {
	mi := &file_send_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessMessage.ProtoReflect.Descriptor instead.
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return file_send_proto_rawDescGZIP(), []int{1}
}

func (x *SuccessMessage) GetIsAdded() string {
	if x != nil {
		return x.IsAdded
	}
	return ""
}

var File_send_proto protoreflect.FileDescriptor

var file_send_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x64,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x73, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x32, 0x51, 0x0a, 0x0c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x41, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_send_proto_rawDescOnce sync.Once
	file_send_proto_rawDescData = file_send_proto_rawDesc
)

func file_send_proto_rawDescGZIP() []byte {
	file_send_proto_rawDescOnce.Do(func() {
		file_send_proto_rawDescData = protoimpl.X.CompressGZIP(file_send_proto_rawDescData)
	})
	return file_send_proto_rawDescData
}

var file_send_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_send_proto_goTypes = []interface{}{
	(*RewardDetails)(nil),  // 0: protos.RewardDetails
	(*SuccessMessage)(nil), // 1: protos.SuccessMessage
}
var file_send_proto_depIdxs = []int32{
	0, // 0: protos.RewardPoints.SendRewardPoints:input_type -> protos.RewardDetails
	1, // 1: protos.RewardPoints.SendRewardPoints:output_type -> protos.SuccessMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_send_proto_init() }
func file_send_proto_init() {
	if File_send_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_send_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardDetails); i {
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
		file_send_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessMessage); i {
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
			RawDescriptor: file_send_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_send_proto_goTypes,
		DependencyIndexes: file_send_proto_depIdxs,
		MessageInfos:      file_send_proto_msgTypes,
	}.Build()
	File_send_proto = out.File
	file_send_proto_rawDesc = nil
	file_send_proto_goTypes = nil
	file_send_proto_depIdxs = nil
}