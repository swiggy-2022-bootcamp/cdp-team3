// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: app/grpcs/Products/products.proto

package products

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

type CategoryDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CategoryDeleteRequest) Reset() {
	*x = CategoryDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpcs_Products_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDeleteRequest) ProtoMessage() {}

func (x *CategoryDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpcs_Products_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDeleteRequest.ProtoReflect.Descriptor instead.
func (*CategoryDeleteRequest) Descriptor() ([]byte, []int) {
	return file_app_grpcs_Products_products_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryDeleteRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type CategoriesDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoriesId []string `protobuf:"bytes,1,rep,name=categories_id,json=categoriesId,proto3" json:"categories_id,omitempty"`
}

func (x *CategoriesDeleteRequest) Reset() {
	*x = CategoriesDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpcs_Products_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesDeleteRequest) ProtoMessage() {}

func (x *CategoriesDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpcs_Products_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesDeleteRequest.ProtoReflect.Descriptor instead.
func (*CategoriesDeleteRequest) Descriptor() ([]byte, []int) {
	return file_app_grpcs_Products_products_proto_rawDescGZIP(), []int{1}
}

func (x *CategoriesDeleteRequest) GetCategoriesId() []string {
	if x != nil {
		return x.CategoriesId
	}
	return nil
}

type DeleteCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *DeleteCategoriesResponse) Reset() {
	*x = DeleteCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpcs_Products_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoriesResponse) ProtoMessage() {}

func (x *DeleteCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpcs_Products_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoriesResponse.ProtoReflect.Descriptor instead.
func (*DeleteCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_app_grpcs_Products_products_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCategoriesResponse) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

type DeleteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *DeleteCategoryResponse) Reset() {
	*x = DeleteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpcs_Products_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryResponse) ProtoMessage() {}

func (x *DeleteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpcs_Products_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) {
	return file_app_grpcs_Products_products_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCategoryResponse) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

var File_app_grpcs_Products_products_proto protoreflect.FileDescriptor

var file_app_grpcs_Products_products_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x38, 0x0a,
	0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x17, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65,
	0x73, 0x32, 0xba, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x59,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11,
	0x5a, 0x0f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_grpcs_Products_products_proto_rawDescOnce sync.Once
	file_app_grpcs_Products_products_proto_rawDescData = file_app_grpcs_Products_products_proto_rawDesc
)

func file_app_grpcs_Products_products_proto_rawDescGZIP() []byte {
	file_app_grpcs_Products_products_proto_rawDescOnce.Do(func() {
		file_app_grpcs_Products_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpcs_Products_products_proto_rawDescData)
	})
	return file_app_grpcs_Products_products_proto_rawDescData
}

var file_app_grpcs_Products_products_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_grpcs_Products_products_proto_goTypes = []interface{}{
	(*CategoryDeleteRequest)(nil),    // 0: products.CategoryDeleteRequest
	(*CategoriesDeleteRequest)(nil),  // 1: products.CategoriesDeleteRequest
	(*DeleteCategoriesResponse)(nil), // 2: products.DeleteCategoriesResponse
	(*DeleteCategoryResponse)(nil),   // 3: products.DeleteCategoryResponse
}
var file_app_grpcs_Products_products_proto_depIdxs = []int32{
	1, // 0: products.Products.DeleteCategories:input_type -> products.CategoriesDeleteRequest
	0, // 1: products.Products.DeleteCategory:input_type -> products.CategoryDeleteRequest
	2, // 2: products.Products.DeleteCategories:output_type -> products.DeleteCategoriesResponse
	3, // 3: products.Products.DeleteCategory:output_type -> products.DeleteCategoryResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_grpcs_Products_products_proto_init() }
func file_app_grpcs_Products_products_proto_init() {
	if File_app_grpcs_Products_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpcs_Products_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDeleteRequest); i {
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
		file_app_grpcs_Products_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesDeleteRequest); i {
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
		file_app_grpcs_Products_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoriesResponse); i {
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
		file_app_grpcs_Products_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryResponse); i {
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
			RawDescriptor: file_app_grpcs_Products_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpcs_Products_products_proto_goTypes,
		DependencyIndexes: file_app_grpcs_Products_products_proto_depIdxs,
		MessageInfos:      file_app_grpcs_Products_products_proto_msgTypes,
	}.Build()
	File_app_grpcs_Products_products_proto = out.File
	file_app_grpcs_Products_products_proto_rawDesc = nil
	file_app_grpcs_Products_products_proto_goTypes = nil
	file_app_grpcs_Products_products_proto_depIdxs = nil
}