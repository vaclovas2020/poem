// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: admin/admin.proto

package admin

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

type AdminCategory_Status int32

const (
	AdminCategory_PUBLISHED AdminCategory_Status = 0
	AdminCategory_DRAFT     AdminCategory_Status = 1
)

// Enum value maps for AdminCategory_Status.
var (
	AdminCategory_Status_name = map[int32]string{
		0: "PUBLISHED",
		1: "DRAFT",
	}
	AdminCategory_Status_value = map[string]int32{
		"PUBLISHED": 0,
		"DRAFT":     1,
	}
)

func (x AdminCategory_Status) Enum() *AdminCategory_Status {
	p := new(AdminCategory_Status)
	*p = x
	return p
}

func (x AdminCategory_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminCategory_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_admin_proto_enumTypes[0].Descriptor()
}

func (AdminCategory_Status) Type() protoreflect.EnumType {
	return &file_admin_admin_proto_enumTypes[0]
}

func (x AdminCategory_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminCategory_Status.Descriptor instead.
func (AdminCategory_Status) EnumDescriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{5, 0}
}

// The response message containing poem  object content.
type PoemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool       `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Poem    *AdminPoem `protobuf:"bytes,2,opt,name=poem,proto3" json:"poem,omitempty"`
}

func (x *PoemResponse) Reset() {
	*x = PoemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoemResponse) ProtoMessage() {}

func (x *PoemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoemResponse.ProtoReflect.Descriptor instead.
func (*PoemResponse) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *PoemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PoemResponse) GetPoem() *AdminPoem {
	if x != nil {
		return x.Poem
	}
	return nil
}

// The response message containing category  object content.
type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Category *AdminCategory `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CategoryResponse) GetCategory() *AdminCategory {
	if x != nil {
		return x.Category
	}
	return nil
}

type DeleteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int32 `protobuf:"varint,1,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	UserId     int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCategoryRequest) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *DeleteCategoryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCategoryResponse) Reset() {
	*x = DeleteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryResponse) ProtoMessage() {}

func (x *DeleteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[3]
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
	return file_admin_admin_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCategoryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AdminPoem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text       string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CategoryId int32  `protobuf:"varint,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	UserId     int64  `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AdminPoem) Reset() {
	*x = AdminPoem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminPoem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminPoem) ProtoMessage() {}

func (x *AdminPoem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminPoem.ProtoReflect.Descriptor instead.
func (*AdminPoem) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AdminPoem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AdminPoem) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AdminPoem) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AdminPoem) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AdminCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug   string               `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Status AdminCategory_Status `protobuf:"varint,3,opt,name=status,proto3,enum=AdminCategory_Status" json:"status,omitempty"`
	UserId int64                `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AdminCategory) Reset() {
	*x = AdminCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCategory) ProtoMessage() {}

func (x *AdminCategory) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCategory.ProtoReflect.Descriptor instead.
func (*AdminCategory) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{5}
}

func (x *AdminCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminCategory) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *AdminCategory) GetStatus() AdminCategory_Status {
	if x != nil {
		return x.Status
	}
	return AdminCategory_PUBLISHED
}

func (x *AdminCategory) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_admin_admin_proto protoreflect.FileDescriptor

var file_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0c, 0x50, 0x6f, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x70, 0x6f, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x65, 0x6d, 0x52, 0x04, 0x70, 0x6f, 0x65, 0x6d, 0x22, 0x58, 0x0a,
	0x10, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x6d, 0x0a, 0x09,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0d,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01,
	0x32, 0xa8, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x65, 0x6d, 0x12, 0x0a, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x65,
	0x6d, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x1a, 0x11, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x77,
	0x65, 0x62, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x6f, 0x65,
	0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_admin_proto_rawDescOnce sync.Once
	file_admin_admin_proto_rawDescData = file_admin_admin_proto_rawDesc
)

func file_admin_admin_proto_rawDescGZIP() []byte {
	file_admin_admin_proto_rawDescOnce.Do(func() {
		file_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_admin_proto_rawDescData)
	})
	return file_admin_admin_proto_rawDescData
}

var file_admin_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_admin_proto_goTypes = []interface{}{
	(AdminCategory_Status)(0),      // 0: AdminCategory.Status
	(*PoemResponse)(nil),           // 1: PoemResponse
	(*CategoryResponse)(nil),       // 2: CategoryResponse
	(*DeleteCategoryRequest)(nil),  // 3: DeleteCategoryRequest
	(*DeleteCategoryResponse)(nil), // 4: DeleteCategoryResponse
	(*AdminPoem)(nil),              // 5: AdminPoem
	(*AdminCategory)(nil),          // 6: AdminCategory
}
var file_admin_admin_proto_depIdxs = []int32{
	5, // 0: PoemResponse.poem:type_name -> AdminPoem
	6, // 1: CategoryResponse.category:type_name -> AdminCategory
	0, // 2: AdminCategory.status:type_name -> AdminCategory.Status
	5, // 3: Admin.AddPoem:input_type -> AdminPoem
	6, // 4: Admin.AddCategory:input_type -> AdminCategory
	3, // 5: Admin.DeleteCategory:input_type -> DeleteCategoryRequest
	1, // 6: Admin.AddPoem:output_type -> PoemResponse
	2, // 7: Admin.AddCategory:output_type -> CategoryResponse
	4, // 8: Admin.DeleteCategory:output_type -> DeleteCategoryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_admin_proto_init() }
func file_admin_admin_proto_init() {
	if File_admin_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoemResponse); i {
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
		file_admin_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_admin_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryRequest); i {
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
		file_admin_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminPoem); i {
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
		file_admin_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCategory); i {
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
			RawDescriptor: file_admin_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_admin_proto_goTypes,
		DependencyIndexes: file_admin_admin_proto_depIdxs,
		EnumInfos:         file_admin_admin_proto_enumTypes,
		MessageInfos:      file_admin_admin_proto_msgTypes,
	}.Build()
	File_admin_admin_proto = out.File
	file_admin_admin_proto_rawDesc = nil
	file_admin_admin_proto_goTypes = nil
	file_admin_admin_proto_depIdxs = nil
}
