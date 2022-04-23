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

// The Signal having the Cart ID or the User ID to do operations with
type CartIDSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartID string `protobuf:"bytes,1,opt,name=cartID,proto3" json:"cartID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *CartIDSignal) Reset() {
	*x = CartIDSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartIDSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartIDSignal) ProtoMessage() {}

func (x *CartIDSignal) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CartIDSignal.ProtoReflect.Descriptor instead.
func (*CartIDSignal) Descriptor() ([]byte, []int) {
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *CartIDSignal) GetCartID() string {
	if x != nil {
		return x.CartID
	}
	return ""
}

func (x *CartIDSignal) GetUserID() string {
	if x != nil {
		return x.UserID
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

// Definition of a Cart Item
type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string  `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *CartItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CartItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

// The Response having the list of Cart Items
type CartItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    bool        `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	CartItems []*CartItem `protobuf:"bytes,2,rep,name=cartItems,proto3" json:"cartItems,omitempty"`
}

func (x *CartItemsResponse) Reset() {
	*x = CartItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemsResponse) ProtoMessage() {}

func (x *CartItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemsResponse.ProtoReflect.Descriptor instead.
func (*CartItemsResponse) Descriptor() ([]byte, []int) {
	return file_grpcs_cart_checkout_cart_checkout_proto_rawDescGZIP(), []int{3}
}

func (x *CartItemsResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CartItemsResponse) GetCartItems() []*CartItem {
	if x != nil {
		return x.CartItems
	}
	return nil
}

var File_grpcs_cart_checkout_cart_checkout_proto protoreflect.FileDescriptor

var file_grpcs_cart_checkout_cart_checkout_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x3e, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x44, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x5a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x62, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x09,
	0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0xb2, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x72, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x44, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_grpcs_cart_checkout_cart_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpcs_cart_checkout_cart_checkout_proto_goTypes = []interface{}{
	(*CartIDSignal)(nil),      // 0: cart_checkout.CartIDSignal
	(*CartEmptyOutput)(nil),   // 1: cart_checkout.CartEmptyOutput
	(*CartItem)(nil),          // 2: cart_checkout.CartItem
	(*CartItemsResponse)(nil), // 3: cart_checkout.CartItemsResponse
}
var file_grpcs_cart_checkout_cart_checkout_proto_depIdxs = []int32{
	2, // 0: cart_checkout.CartItemsResponse.cartItems:type_name -> cart_checkout.CartItem
	0, // 1: cart_checkout.CartCheckoutService.GetCartItems:input_type -> cart_checkout.CartIDSignal
	0, // 2: cart_checkout.CartCheckoutService.EmptyCart:input_type -> cart_checkout.CartIDSignal
	3, // 3: cart_checkout.CartCheckoutService.GetCartItems:output_type -> cart_checkout.CartItemsResponse
	1, // 4: cart_checkout.CartCheckoutService.EmptyCart:output_type -> cart_checkout.CartEmptyOutput
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpcs_cart_checkout_cart_checkout_proto_init() }
func file_grpcs_cart_checkout_cart_checkout_proto_init() {
	if File_grpcs_cart_checkout_cart_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartIDSignal); i {
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
		file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_grpcs_cart_checkout_cart_checkout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemsResponse); i {
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
			NumMessages:   4,
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