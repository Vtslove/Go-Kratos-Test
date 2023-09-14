// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: content/service/v1/tag.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-cms/gen/api/go/common/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 标签
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Slug       *string `protobuf:"bytes,3,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	SlugName   *string `protobuf:"bytes,4,opt,name=slugName,proto3,oneof" json:"slugName,omitempty"`
	Color      *string `protobuf:"bytes,5,opt,name=color,proto3,oneof" json:"color,omitempty"`
	Thumbnail  *string `protobuf:"bytes,6,opt,name=thumbnail,proto3,oneof" json:"thumbnail,omitempty"`
	CreateTime *string `protobuf:"bytes,7,opt,name=createTime,proto3,oneof" json:"createTime,omitempty"`
	UpdateTime *string `protobuf:"bytes,8,opt,name=updateTime,proto3,oneof" json:"updateTime,omitempty"`
	PostCount  *uint32 `protobuf:"varint,9,opt,name=postCount,proto3,oneof" json:"postCount,omitempty"`
	FullPath   *string `protobuf:"bytes,10,opt,name=fullPath,proto3,oneof" json:"fullPath,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Tag) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *Tag) GetSlugName() string {
	if x != nil && x.SlugName != nil {
		return *x.SlugName
	}
	return ""
}

func (x *Tag) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *Tag) GetThumbnail() string {
	if x != nil && x.Thumbnail != nil {
		return *x.Thumbnail
	}
	return ""
}

func (x *Tag) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

func (x *Tag) GetUpdateTime() string {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return ""
}

func (x *Tag) GetPostCount() uint32 {
	if x != nil && x.PostCount != nil {
		return *x.PostCount
	}
	return 0
}

func (x *Tag) GetFullPath() string {
	if x != nil && x.FullPath != nil {
		return *x.FullPath
	}
	return ""
}

// 回应 - 标签列表
type ListTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Tag `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListTagResponse) Reset() {
	*x = ListTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagResponse) ProtoMessage() {}

func (x *ListTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagResponse.ProtoReflect.Descriptor instead.
func (*ListTagResponse) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{1}
}

func (x *ListTagResponse) GetItems() []*Tag {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTagResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 标签数据
type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求 - 创建标签
type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag        *Tag    `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *CreateTagRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 更新标签
type UpdateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tag        *Tag    `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	OperatorId *uint32 `protobuf:"varint,3,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdateTagRequest) Reset() {
	*x = UpdateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagRequest) ProtoMessage() {}

func (x *UpdateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateTagRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTagRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *UpdateTagRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 删除标签
type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_tag_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTagRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTagRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_content_service_v1_tag_proto protoreflect.FileDescriptor

var file_content_service_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x73, 0x6c, 0x75, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x08, 0x73, 0x6c, 0x75, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x05, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x73, 0x6c, 0x75, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x22, 0x56, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x32, 0x8a, 0x03, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x19, 0x2e, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12,
	0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x2d, 0x5a, 0x2b, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_service_v1_tag_proto_rawDescOnce sync.Once
	file_content_service_v1_tag_proto_rawDescData = file_content_service_v1_tag_proto_rawDesc
)

func file_content_service_v1_tag_proto_rawDescGZIP() []byte {
	file_content_service_v1_tag_proto_rawDescOnce.Do(func() {
		file_content_service_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_service_v1_tag_proto_rawDescData)
	})
	return file_content_service_v1_tag_proto_rawDescData
}

var file_content_service_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_content_service_v1_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),                      // 0: content.service.v1.Tag
	(*ListTagResponse)(nil),          // 1: content.service.v1.ListTagResponse
	(*GetTagRequest)(nil),            // 2: content.service.v1.GetTagRequest
	(*CreateTagRequest)(nil),         // 3: content.service.v1.CreateTagRequest
	(*UpdateTagRequest)(nil),         // 4: content.service.v1.UpdateTagRequest
	(*DeleteTagRequest)(nil),         // 5: content.service.v1.DeleteTagRequest
	(*pagination.PagingRequest)(nil), // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_content_service_v1_tag_proto_depIdxs = []int32{
	0, // 0: content.service.v1.ListTagResponse.items:type_name -> content.service.v1.Tag
	0, // 1: content.service.v1.CreateTagRequest.tag:type_name -> content.service.v1.Tag
	0, // 2: content.service.v1.UpdateTagRequest.tag:type_name -> content.service.v1.Tag
	6, // 3: content.service.v1.TagService.ListTag:input_type -> pagination.PagingRequest
	2, // 4: content.service.v1.TagService.GetTag:input_type -> content.service.v1.GetTagRequest
	3, // 5: content.service.v1.TagService.CreateTag:input_type -> content.service.v1.CreateTagRequest
	4, // 6: content.service.v1.TagService.UpdateTag:input_type -> content.service.v1.UpdateTagRequest
	5, // 7: content.service.v1.TagService.DeleteTag:input_type -> content.service.v1.DeleteTagRequest
	1, // 8: content.service.v1.TagService.ListTag:output_type -> content.service.v1.ListTagResponse
	0, // 9: content.service.v1.TagService.GetTag:output_type -> content.service.v1.Tag
	0, // 10: content.service.v1.TagService.CreateTag:output_type -> content.service.v1.Tag
	0, // 11: content.service.v1.TagService.UpdateTag:output_type -> content.service.v1.Tag
	7, // 12: content.service.v1.TagService.DeleteTag:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_content_service_v1_tag_proto_init() }
func file_content_service_v1_tag_proto_init() {
	if File_content_service_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_service_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_content_service_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagResponse); i {
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
		file_content_service_v1_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_content_service_v1_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_content_service_v1_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagRequest); i {
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
		file_content_service_v1_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
	file_content_service_v1_tag_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_content_service_v1_tag_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_content_service_v1_tag_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_content_service_v1_tag_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_service_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_service_v1_tag_proto_goTypes,
		DependencyIndexes: file_content_service_v1_tag_proto_depIdxs,
		MessageInfos:      file_content_service_v1_tag_proto_msgTypes,
	}.Build()
	File_content_service_v1_tag_proto = out.File
	file_content_service_v1_tag_proto_rawDesc = nil
	file_content_service_v1_tag_proto_goTypes = nil
	file_content_service_v1_tag_proto_depIdxs = nil
}
