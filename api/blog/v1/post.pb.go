// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: post.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/third_party/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           *string `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Status          *int32  `protobuf:"varint,3,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Slug            *string `protobuf:"bytes,4,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	EditorType      *int32  `protobuf:"varint,5,opt,name=editorType,proto3,oneof" json:"editorType,omitempty"`
	MetaKeywords    *string `protobuf:"bytes,6,opt,name=metaKeywords,proto3,oneof" json:"metaKeywords,omitempty"`
	MetaDescription *string `protobuf:"bytes,7,opt,name=metaDescription,proto3,oneof" json:"metaDescription,omitempty"`
	FullPath        *string `protobuf:"bytes,8,opt,name=fullPath,proto3,oneof" json:"fullPath,omitempty"`
	Summary         *string `protobuf:"bytes,9,opt,name=summary,proto3,oneof" json:"summary,omitempty"`
	Thumbnail       *string `protobuf:"bytes,10,opt,name=thumbnail,proto3,oneof" json:"thumbnail,omitempty"`
	Password        *string `protobuf:"bytes,11,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Template        *string `protobuf:"bytes,12,opt,name=template,proto3,oneof" json:"template,omitempty"`
	Content         *string `protobuf:"bytes,13,opt,name=content,proto3,oneof" json:"content,omitempty"`
	OriginalContent *string `protobuf:"bytes,14,opt,name=originalContent,proto3,oneof" json:"originalContent,omitempty"`
	Visits          *int32  `protobuf:"varint,15,opt,name=visits,proto3,oneof" json:"visits,omitempty"`
	TopPriority     *int32  `protobuf:"varint,16,opt,name=topPriority,proto3,oneof" json:"topPriority,omitempty"`
	Likes           *int32  `protobuf:"varint,17,opt,name=likes,proto3,oneof" json:"likes,omitempty"`
	WordCount       *int32  `protobuf:"varint,18,opt,name=wordCount,proto3,oneof" json:"wordCount,omitempty"`
	CommentCount    *int32  `protobuf:"varint,19,opt,name=commentCount,proto3,oneof" json:"commentCount,omitempty"`
	DisallowComment *bool   `protobuf:"varint,20,opt,name=disallowComment,proto3,oneof" json:"disallowComment,omitempty"`
	InProgress      *bool   `protobuf:"varint,21,opt,name=inProgress,proto3,oneof" json:"inProgress,omitempty"`
	CreateTime      *string `protobuf:"bytes,22,opt,name=createTime,proto3,oneof" json:"createTime,omitempty"`
	UpdateTime      *string `protobuf:"bytes,23,opt,name=updateTime,proto3,oneof" json:"updateTime,omitempty"`
	EditTime        *string `protobuf:"bytes,24,opt,name=editTime,proto3,oneof" json:"editTime,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Post) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Post) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *Post) GetEditorType() int32 {
	if x != nil && x.EditorType != nil {
		return *x.EditorType
	}
	return 0
}

func (x *Post) GetMetaKeywords() string {
	if x != nil && x.MetaKeywords != nil {
		return *x.MetaKeywords
	}
	return ""
}

func (x *Post) GetMetaDescription() string {
	if x != nil && x.MetaDescription != nil {
		return *x.MetaDescription
	}
	return ""
}

func (x *Post) GetFullPath() string {
	if x != nil && x.FullPath != nil {
		return *x.FullPath
	}
	return ""
}

func (x *Post) GetSummary() string {
	if x != nil && x.Summary != nil {
		return *x.Summary
	}
	return ""
}

func (x *Post) GetThumbnail() string {
	if x != nil && x.Thumbnail != nil {
		return *x.Thumbnail
	}
	return ""
}

func (x *Post) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *Post) GetTemplate() string {
	if x != nil && x.Template != nil {
		return *x.Template
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *Post) GetOriginalContent() string {
	if x != nil && x.OriginalContent != nil {
		return *x.OriginalContent
	}
	return ""
}

func (x *Post) GetVisits() int32 {
	if x != nil && x.Visits != nil {
		return *x.Visits
	}
	return 0
}

func (x *Post) GetTopPriority() int32 {
	if x != nil && x.TopPriority != nil {
		return *x.TopPriority
	}
	return 0
}

func (x *Post) GetLikes() int32 {
	if x != nil && x.Likes != nil {
		return *x.Likes
	}
	return 0
}

func (x *Post) GetWordCount() int32 {
	if x != nil && x.WordCount != nil {
		return *x.WordCount
	}
	return 0
}

func (x *Post) GetCommentCount() int32 {
	if x != nil && x.CommentCount != nil {
		return *x.CommentCount
	}
	return 0
}

func (x *Post) GetDisallowComment() bool {
	if x != nil && x.DisallowComment != nil {
		return *x.DisallowComment
	}
	return false
}

func (x *Post) GetInProgress() bool {
	if x != nil && x.InProgress != nil {
		return *x.InProgress
	}
	return false
}

func (x *Post) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

func (x *Post) GetUpdateTime() string {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return ""
}

func (x *Post) GetEditTime() string {
	if x != nil && x.EditTime != nil {
		return *x.EditTime
	}
	return ""
}

// 回应 - 帖子列表
type ListPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Post `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListPostResponse) Reset() {
	*x = ListPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostResponse) ProtoMessage() {}

func (x *ListPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostResponse.ProtoReflect.Descriptor instead.
func (*ListPostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *ListPostResponse) GetItems() []*Post {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPostResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 帖子数据
type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求 - 创建帖子
type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post       *Post   `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *CreatePostRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 更新帖子
type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Post       *Post   `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	OperatorId *uint32 `protobuf:"varint,3,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePostRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *UpdatePostRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 删除帖子
type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeletePostRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x09,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0a, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x6d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x0c, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0d, 0x52, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0e, 0x52, 0x0b, 0x74, 0x6f, 0x70,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0f, 0x52, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x48, 0x10, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x11, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x48, 0x12, 0x52, 0x0f, 0x64,
	0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x08, 0x48, 0x13, 0x52, 0x0a, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x48, 0x14, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x15, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x65, 0x64, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x16, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x69, 0x73,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x64,
	0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x57, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xc6, 0x03, 0x0a, 0x0b, 0x50, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x4e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x52, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x13, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x12, 0x5d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x1c, 0x5a, 0x1a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_post_proto_goTypes = []interface{}{
	(*Post)(nil),                     // 0: blog.v1.Post
	(*ListPostResponse)(nil),         // 1: blog.v1.ListPostResponse
	(*GetPostRequest)(nil),           // 2: blog.v1.GetPostRequest
	(*CreatePostRequest)(nil),        // 3: blog.v1.CreatePostRequest
	(*UpdatePostRequest)(nil),        // 4: blog.v1.UpdatePostRequest
	(*DeletePostRequest)(nil),        // 5: blog.v1.DeletePostRequest
	(*pagination.PagingRequest)(nil), // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_post_proto_depIdxs = []int32{
	0, // 0: blog.v1.ListPostResponse.items:type_name -> blog.v1.Post
	0, // 1: blog.v1.CreatePostRequest.post:type_name -> blog.v1.Post
	0, // 2: blog.v1.UpdatePostRequest.post:type_name -> blog.v1.Post
	6, // 3: blog.v1.PostService.ListPost:input_type -> pagination.PagingRequest
	2, // 4: blog.v1.PostService.GetPost:input_type -> blog.v1.GetPostRequest
	3, // 5: blog.v1.PostService.CreatePost:input_type -> blog.v1.CreatePostRequest
	4, // 6: blog.v1.PostService.UpdatePost:input_type -> blog.v1.UpdatePostRequest
	5, // 7: blog.v1.PostService.DeletePost:input_type -> blog.v1.DeletePostRequest
	1, // 8: blog.v1.PostService.ListPost:output_type -> blog.v1.ListPostResponse
	0, // 9: blog.v1.PostService.GetPost:output_type -> blog.v1.Post
	0, // 10: blog.v1.PostService.CreatePost:output_type -> blog.v1.Post
	0, // 11: blog.v1.PostService.UpdatePost:output_type -> blog.v1.Post
	7, // 12: blog.v1.PostService.DeletePost:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostResponse); i {
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
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
	file_post_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_post_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_post_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_post_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
