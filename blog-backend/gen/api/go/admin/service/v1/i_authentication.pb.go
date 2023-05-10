// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: admin/service/v1/i_authentication.proto

package v1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/gen/api/go/user/service/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 用户后台登陆 - 请求
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName *string `protobuf:"bytes,1,opt,name=userName,proto3,oneof" json:"userName,omitempty"`
	Password *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

// 用户后台登陆 - 回应
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token        string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// 用户后台登出 - 请求
type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取当前用户信息 - 请求
type GetMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMeRequest) Reset() {
	*x = GetMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeRequest) ProtoMessage() {}

func (x *GetMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeRequest.ProtoReflect.Descriptor instead.
func (*GetMeRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *GetMeRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_admin_service_v1_i_authentication_proto protoreflect.FileDescriptor

var file_admin_service_v1_i_authentication_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x75, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb3, 0x02, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x5e, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x54, 0x0a, 0x05,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x42, 0x70, 0x5a, 0x2a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0xba, 0x47, 0x41, 0x12, 0x3f, 0x0a, 0x1e, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0x90, 0x8e,
	0xe5, 0x8f, 0xb0, 0xe7, 0x99, 0xbb, 0xe9, 0x99, 0x86, 0xe8, 0xae, 0xa4, 0xe8, 0xaf, 0x81, 0xe6,
	0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x12, 0x18, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0x90, 0x8e,
	0xe5, 0x8f, 0xb0, 0xe7, 0x99, 0xbb, 0xe9, 0x99, 0x86, 0xe8, 0xae, 0xa4, 0xe8, 0xaf, 0x81, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_service_v1_i_authentication_proto_rawDescOnce sync.Once
	file_admin_service_v1_i_authentication_proto_rawDescData = file_admin_service_v1_i_authentication_proto_rawDesc
)

func file_admin_service_v1_i_authentication_proto_rawDescGZIP() []byte {
	file_admin_service_v1_i_authentication_proto_rawDescOnce.Do(func() {
		file_admin_service_v1_i_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_service_v1_i_authentication_proto_rawDescData)
	})
	return file_admin_service_v1_i_authentication_proto_rawDescData
}

var file_admin_service_v1_i_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_service_v1_i_authentication_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),  // 0: admin.service.v1.LoginRequest
	(*LoginResponse)(nil), // 1: admin.service.v1.LoginResponse
	(*LogoutRequest)(nil), // 2: admin.service.v1.LogoutRequest
	(*GetMeRequest)(nil),  // 3: admin.service.v1.GetMeRequest
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
	(*v1.User)(nil),       // 5: user.service.v1.User
}
var file_admin_service_v1_i_authentication_proto_depIdxs = []int32{
	0, // 0: admin.service.v1.AuthenticationService.Login:input_type -> admin.service.v1.LoginRequest
	2, // 1: admin.service.v1.AuthenticationService.Logout:input_type -> admin.service.v1.LogoutRequest
	3, // 2: admin.service.v1.AuthenticationService.GetMe:input_type -> admin.service.v1.GetMeRequest
	1, // 3: admin.service.v1.AuthenticationService.Login:output_type -> admin.service.v1.LoginResponse
	4, // 4: admin.service.v1.AuthenticationService.Logout:output_type -> google.protobuf.Empty
	5, // 5: admin.service.v1.AuthenticationService.GetMe:output_type -> user.service.v1.User
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_i_authentication_proto_init() }
func file_admin_service_v1_i_authentication_proto_init() {
	if File_admin_service_v1_i_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_service_v1_i_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_admin_service_v1_i_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_admin_service_v1_i_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_admin_service_v1_i_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeRequest); i {
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
	file_admin_service_v1_i_authentication_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_i_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_v1_i_authentication_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_i_authentication_proto_depIdxs,
		MessageInfos:      file_admin_service_v1_i_authentication_proto_msgTypes,
	}.Build()
	File_admin_service_v1_i_authentication_proto = out.File
	file_admin_service_v1_i_authentication_proto_rawDesc = nil
	file_admin_service_v1_i_authentication_proto_goTypes = nil
	file_admin_service_v1_i_authentication_proto_depIdxs = nil
}
