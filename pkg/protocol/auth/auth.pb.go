// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: auth/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret"`
	PlatformID int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID"`
	UserID     string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID"`
}

func (x *UserTokenReq) Reset() {
	*x = UserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenReq) ProtoMessage() {}

func (x *UserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenReq.ProtoReflect.Descriptor instead.
func (*UserTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserTokenReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *UserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	ExpireTimeSeconds int64  `protobuf:"varint,3,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
}

func (x *UserTokenResp) Reset() {
	*x = UserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenResp) ProtoMessage() {}

func (x *UserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenResp.ProtoReflect.Descriptor instead.
func (*UserTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type ForceLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
}

func (x *ForceLogoutReq) Reset() {
	*x = ForceLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutReq) ProtoMessage() {}

func (x *ForceLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutReq.ProtoReflect.Descriptor instead.
func (*ForceLogoutReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ForceLogoutReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ForceLogoutReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ForceLogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceLogoutResp) Reset() {
	*x = ForceLogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutResp) ProtoMessage() {}

func (x *ForceLogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutResp.ProtoReflect.Descriptor instead.
func (*ForceLogoutResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

type ParseTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
}

func (x *ParseTokenReq) Reset() {
	*x = ParseTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenReq) ProtoMessage() {}

func (x *ParseTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenReq.ProtoReflect.Descriptor instead.
func (*ParseTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ParseTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID            string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
	PlatformID        int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID"`
	ExpireTimeSeconds int64  `protobuf:"varint,4,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
}

func (x *ParseTokenResp) Reset() {
	*x = ParseTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResp) ProtoMessage() {}

func (x *ParseTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResp.ProtoReflect.Descriptor instead.
func (*ParseTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ParseTokenResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParseTokenResp) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ParseTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *GetUserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	ExpireTimeSeconds int64  `protobuf:"varint,2,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
}

func (x *GetUserTokenResp) Reset() {
	*x = GetUserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResp) ProtoMessage() {}

func (x *GetUserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResp.ProtoReflect.Descriptor instead.
func (*GetUserTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type InvalidateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreservedToken string `protobuf:"bytes,1,opt,name=preservedToken,proto3" json:"preservedToken"`
	UserID         string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
	PlatformID     int32  `protobuf:"varint,3,opt,name=platformID,proto3" json:"platformID"`
}

func (x *InvalidateTokenReq) Reset() {
	*x = InvalidateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateTokenReq) ProtoMessage() {}

func (x *InvalidateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateTokenReq.ProtoReflect.Descriptor instead.
func (*InvalidateTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *InvalidateTokenReq) GetPreservedToken() string {
	if x != nil {
		return x.PreservedToken
	}
	return ""
}

func (x *InvalidateTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *InvalidateTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

type InvalidateTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InvalidateTokenResp) Reset() {
	*x = InvalidateTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateTokenResp) ProtoMessage() {}

func (x *InvalidateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateTokenResp.ProtoReflect.Descriptor instead.
func (*InvalidateTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{9}
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x5e,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x53,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0x48, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x11, 0x0a,
	0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x25, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x76, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x44, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22,
	0x49, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x56, 0x0a, 0x10, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x22, 0x74, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x22, 0x15, 0x0a, 0x13, 0x69, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x32,
	0xfe, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x42, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0c,
	0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0b, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0f, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x69, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x69, 0x6d,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_auth_auth_proto_goTypes = []interface{}{
	(*UserTokenReq)(nil),        // 0: openim.auth.userTokenReq
	(*UserTokenResp)(nil),       // 1: openim.auth.userTokenResp
	(*ForceLogoutReq)(nil),      // 2: openim.auth.forceLogoutReq
	(*ForceLogoutResp)(nil),     // 3: openim.auth.forceLogoutResp
	(*ParseTokenReq)(nil),       // 4: openim.auth.parseTokenReq
	(*ParseTokenResp)(nil),      // 5: openim.auth.parseTokenResp
	(*GetUserTokenReq)(nil),     // 6: openim.auth.getUserTokenReq
	(*GetUserTokenResp)(nil),    // 7: openim.auth.getUserTokenResp
	(*InvalidateTokenReq)(nil),  // 8: openim.auth.invalidateTokenReq
	(*InvalidateTokenResp)(nil), // 9: openim.auth.invalidateTokenResp
}
var file_auth_auth_proto_depIdxs = []int32{
	0, // 0: openim.auth.Auth.userToken:input_type -> openim.auth.userTokenReq
	6, // 1: openim.auth.Auth.getUserToken:input_type -> openim.auth.getUserTokenReq
	2, // 2: openim.auth.Auth.forceLogout:input_type -> openim.auth.forceLogoutReq
	4, // 3: openim.auth.Auth.parseToken:input_type -> openim.auth.parseTokenReq
	8, // 4: openim.auth.Auth.invalidateToken:input_type -> openim.auth.invalidateTokenReq
	1, // 5: openim.auth.Auth.userToken:output_type -> openim.auth.userTokenResp
	7, // 6: openim.auth.Auth.getUserToken:output_type -> openim.auth.getUserTokenResp
	3, // 7: openim.auth.Auth.forceLogout:output_type -> openim.auth.forceLogoutResp
	5, // 8: openim.auth.Auth.parseToken:output_type -> openim.auth.parseTokenResp
	9, // 9: openim.auth.Auth.invalidateToken:output_type -> openim.auth.invalidateTokenResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenReq); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenResp); i {
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
		file_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutReq); i {
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
		file_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutResp); i {
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
		file_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenReq); i {
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
		file_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenResp); i {
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
		file_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReq); i {
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
		file_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenResp); i {
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
		file_auth_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidateTokenReq); i {
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
		file_auth_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidateTokenResp); i {
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
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// Generate token
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
	// Admin retrieves user token
	GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error)
	// Force logout
	ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error)
	// Parse token
	ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error)
	// Invalidate or mark the token as kicked out
	InvalidateToken(ctx context.Context, in *InvalidateTokenReq, opts ...grpc.CallOption) (*InvalidateTokenResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := c.cc.Invoke(ctx, "/openim.auth.Auth/userToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error) {
	out := new(GetUserTokenResp)
	err := c.cc.Invoke(ctx, "/openim.auth.Auth/getUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error) {
	out := new(ForceLogoutResp)
	err := c.cc.Invoke(ctx, "/openim.auth.Auth/forceLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error) {
	out := new(ParseTokenResp)
	err := c.cc.Invoke(ctx, "/openim.auth.Auth/parseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) InvalidateToken(ctx context.Context, in *InvalidateTokenReq, opts ...grpc.CallOption) (*InvalidateTokenResp, error) {
	out := new(InvalidateTokenResp)
	err := c.cc.Invoke(ctx, "/openim.auth.Auth/invalidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// Generate token
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
	// Admin retrieves user token
	GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error)
	// Force logout
	ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error)
	// Parse token
	ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error)
	// Invalidate or mark the token as kicked out
	InvalidateToken(context.Context, *InvalidateTokenReq) (*InvalidateTokenResp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}
func (*UnimplementedAuthServer) GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserToken not implemented")
}
func (*UnimplementedAuthServer) ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceLogout not implemented")
}
func (*UnimplementedAuthServer) ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseToken not implemented")
}
func (*UnimplementedAuthServer) InvalidateToken(context.Context, *InvalidateTokenReq) (*InvalidateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.auth.Auth/UserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.auth.Auth/GetUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserToken(ctx, req.(*GetUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ForceLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceLogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ForceLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.auth.Auth/ForceLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ForceLogout(ctx, req.(*ForceLogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ParseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ParseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.auth.Auth/ParseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ParseToken(ctx, req.(*ParseTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_InvalidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).InvalidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.auth.Auth/InvalidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).InvalidateToken(ctx, req.(*InvalidateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openim.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userToken",
			Handler:    _Auth_UserToken_Handler,
		},
		{
			MethodName: "getUserToken",
			Handler:    _Auth_GetUserToken_Handler,
		},
		{
			MethodName: "forceLogout",
			Handler:    _Auth_ForceLogout_Handler,
		},
		{
			MethodName: "parseToken",
			Handler:    _Auth_ParseToken_Handler,
		},
		{
			MethodName: "invalidateToken",
			Handler:    _Auth_InvalidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}