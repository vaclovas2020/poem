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

type AdminPoem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AdminPoem) Reset() {
	*x = AdminPoem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminPoem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminPoem) ProtoMessage() {}

func (x *AdminPoem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AdminPoem.ProtoReflect.Descriptor instead.
func (*AdminPoem) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{2}
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

type AdminCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *AdminCategory) Reset() {
	*x = AdminCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCategory) ProtoMessage() {}

func (x *AdminCategory) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AdminCategory.ProtoReflect.Descriptor instead.
func (*AdminCategory) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{3}
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
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x35, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x6f, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x37,
	0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x32, 0x63, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x26, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x65, 0x6d, 0x12, 0x0a, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x65, 0x6d, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18,
	0x77, 0x65, 0x62, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x6f,
	0x65, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_admin_proto_goTypes = []interface{}{
	(*PoemResponse)(nil),     // 0: PoemResponse
	(*CategoryResponse)(nil), // 1: CategoryResponse
	(*AdminPoem)(nil),        // 2: AdminPoem
	(*AdminCategory)(nil),    // 3: AdminCategory
}
var file_admin_admin_proto_depIdxs = []int32{
	2, // 0: PoemResponse.poem:type_name -> AdminPoem
	3, // 1: CategoryResponse.category:type_name -> AdminCategory
	2, // 2: Admin.AddPoem:input_type -> AdminPoem
	3, // 3: Admin.AddCategory:input_type -> AdminCategory
	0, // 4: Admin.AddPoem:output_type -> PoemResponse
	1, // 5: Admin.AddCategory:output_type -> CategoryResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_admin_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_admin_proto_goTypes,
		DependencyIndexes: file_admin_admin_proto_depIdxs,
		MessageInfos:      file_admin_admin_proto_msgTypes,
	}.Build()
	File_admin_admin_proto = out.File
	file_admin_admin_proto_rawDesc = nil
	file_admin_admin_proto_goTypes = nil
	file_admin_admin_proto_depIdxs = nil
}
