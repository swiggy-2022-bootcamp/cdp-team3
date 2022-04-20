// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: grpcs/cart_checkout/cart_checkout.proto

package cart_checkout

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

// The Signal having the cart ID to clear.
type CartEmptySignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartID string `protobuf:"bytes,1,opt,name=cartID,proto3" json:"cartID,omitempty"`
}

func (x *CartEmptySignal) Reset() {
	*x = CartEmptySignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartEmptySignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartEmptySignal) ProtoMessage() {}

func (x *CartEmptySignal) ProtoReflect() protoreflect.Message {
	mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartEmptySignal.ProtoReflect.Descriptor instead.
func (*CartEmptySignal) Descriptor() ([]byte, []int) {
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *CartEmptySignal) GetCartID() string {
	if x != nil {
		return x.CartID
	}
	return ""
}

// The Signal having the boolean value of whether the cart was cleared.
type CartEmptyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CartEmptyOutput) Reset() {
	*x = CartEmptyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartEmptyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartEmptyOutput) ProtoMessage() {}

func (x *CartEmptyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartEmptyOutput.ProtoReflect.Descriptor instead.
func (*CartEmptyOutput) Descriptor() ([]byte, []int) {
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *CartEmptyOutput) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_grpcs_cart_checkout_cart_checkout_proto protoreflect.FileDescriptor

var file_grpcs_cart_checkout_cart_checkout_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x74, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x64,
	0x0a, 0x13, 0x43, 0x61, 0x72, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcs_cart_checkout_cart_checkout_proto_rawDescOnce sync.Once
	file_grpcs_cart_checkout_cart_checkout_proto_rawDescData = file_grpcs_cart_checkout_cart_checkout_proto_rawDesc
)

func file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP() []byte {
	file_grpcs_cart_checkout_cart_checkout_proto_rawDescOnce.Do(func() {
		file_grpcs_cart_checkout_cart_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcs_cart_checkout_cart_checkout_proto_rawDescData)
	})
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescData
}

var file_grpcs_cart_checkout_cart_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcs_cart_checkout_cart_checkout_proto_goTypes = []interface{}{
	(*CartEmptySignal)(nil), // 0: cart_checkout.CartEmptySignal
	(*CartEmptyOutput)(nil), // 1: cart_checkout.CartEmptyOutput
}
var file_grpcs_cart_checkout_cart_checkout_proto_depIdxs = []int32{
	0, // 0: cart_checkout.CartCheckoutService.EmptyCart:input_type -> cart_checkout.CartEmptySignal
	1, // 1: cart_checkout.CartCheckoutService.EmptyCart:output_type -> cart_checkout.CartEmptyOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcs_cart_checkout_cart_checkout_proto_init() }
func file_grpcs_cart_checkout_cart_checkout_proto_init() {
	if File_grpcs_cart_checkout_cart_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartEmptySignal); i {
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
		file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartEmptyOutput); i {
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
			RawDescriptor: file_grpcs_cart_checkout_cart_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcs_cart_checkout_cart_checkout_proto_goTypes,
		DependencyIndexes: file_grpcs_cart_checkout_cart_checkout_proto_depIdxs,
		MessageInfos:      file_grpcs_cart_checkout_cart_checkout_proto_msgTypes,
	}.Build()
	File_grpcs_cart_checkout_cart_checkout_proto = out.File
	file_grpcs_cart_checkout_cart_checkout_proto_rawDesc = nil
	file_grpcs_cart_checkout_cart_checkout_proto_goTypes = nil
	file_grpcs_cart_checkout_cart_checkout_proto_depIdxs = nil
}
