// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v1/captcha.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetSmsCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetSmsCodeReq) Reset() {
	*x = GetSmsCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsCodeReq) ProtoMessage() {}

func (x *GetSmsCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsCodeReq.ProtoReflect.Descriptor instead.
func (*GetSmsCodeReq) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *GetSmsCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetSmsCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetSmsCodeReply) Reset() {
	*x = GetSmsCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsCodeReply) ProtoMessage() {}

func (x *GetSmsCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsCodeReply.ProtoReflect.Descriptor instead.
func (*GetSmsCodeReply) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *GetSmsCodeReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SendSmsCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendSmsCodeReq) Reset() {
	*x = SendSmsCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeReq) ProtoMessage() {}

func (x *SendSmsCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeReq.ProtoReflect.Descriptor instead.
func (*SendSmsCodeReq) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *SendSmsCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SendSmsCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSmsCodeReply) Reset() {
	*x = SendSmsCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeReply) ProtoMessage() {}

func (x *SendSmsCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeReply.ProtoReflect.Descriptor instead.
func (*SendSmsCodeReply) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{3}
}

type GetImageCodeFromRdbReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetImageCodeFromRdbReq) Reset() {
	*x = GetImageCodeFromRdbReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageCodeFromRdbReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageCodeFromRdbReq) ProtoMessage() {}

func (x *GetImageCodeFromRdbReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageCodeFromRdbReq.ProtoReflect.Descriptor instead.
func (*GetImageCodeFromRdbReq) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{4}
}

func (x *GetImageCodeFromRdbReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetImageCodeFromRdbReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImgCode string `protobuf:"bytes,1,opt,name=ImgCode,proto3" json:"ImgCode,omitempty"`
}

func (x *GetImageCodeFromRdbReply) Reset() {
	*x = GetImageCodeFromRdbReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageCodeFromRdbReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageCodeFromRdbReply) ProtoMessage() {}

func (x *GetImageCodeFromRdbReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageCodeFromRdbReply.ProtoReflect.Descriptor instead.
func (*GetImageCodeFromRdbReply) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{5}
}

func (x *GetImageCodeFromRdbReply) GetImgCode() string {
	if x != nil {
		return x.ImgCode
	}
	return ""
}

type GetCaptchaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetCaptchaReq) Reset() {
	*x = GetCaptchaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaptchaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaptchaReq) ProtoMessage() {}

func (x *GetCaptchaReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaptchaReq.ProtoReflect.Descriptor instead.
func (*GetCaptchaReq) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{6}
}

func (x *GetCaptchaReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetCaptchaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ????????????????????????????????? ???json ?????????
	Img []byte `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *GetCaptchaReply) Reset() {
	*x = GetCaptchaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_captcha_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaptchaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaptchaReply) ProtoMessage() {}

func (x *GetCaptchaReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_captcha_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaptchaReply.ProtoReflect.Descriptor instead.
func (*GetCaptchaReply) Descriptor() ([]byte, []int) {
	return file_v1_captcha_proto_rawDescGZIP(), []int{7}
}

func (x *GetCaptchaReply) GetImg() []byte {
	if x != nil {
		return x.Img
	}
	return nil
}

var File_v1_captcha_proto protoreflect.FileDescriptor

var file_v1_captcha_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x52, 0x64, 0x62, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x52, 0x64, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x32, 0x87, 0x03, 0x0a,
	0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x52, 0x64, 0x62, 0x12, 0x2a, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x64, 0x62,
	0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x64, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_captcha_proto_rawDescOnce sync.Once
	file_v1_captcha_proto_rawDescData = file_v1_captcha_proto_rawDesc
)

func file_v1_captcha_proto_rawDescGZIP() []byte {
	file_v1_captcha_proto_rawDescOnce.Do(func() {
		file_v1_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_captcha_proto_rawDescData)
	})
	return file_v1_captcha_proto_rawDescData
}

var file_v1_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_captcha_proto_goTypes = []interface{}{
	(*GetSmsCodeReq)(nil),            // 0: captcha.service.v1.GetSmsCodeReq
	(*GetSmsCodeReply)(nil),          // 1: captcha.service.v1.GetSmsCodeReply
	(*SendSmsCodeReq)(nil),           // 2: captcha.service.v1.SendSmsCodeReq
	(*SendSmsCodeReply)(nil),         // 3: captcha.service.v1.SendSmsCodeReply
	(*GetImageCodeFromRdbReq)(nil),   // 4: captcha.service.v1.GetImageCodeFromRdbReq
	(*GetImageCodeFromRdbReply)(nil), // 5: captcha.service.v1.GetImageCodeFromRdbReply
	(*GetCaptchaReq)(nil),            // 6: captcha.service.v1.GetCaptchaReq
	(*GetCaptchaReply)(nil),          // 7: captcha.service.v1.GetCaptchaReply
}
var file_v1_captcha_proto_depIdxs = []int32{
	6, // 0: captcha.service.v1.Captcha.GetCaptcha:input_type -> captcha.service.v1.GetCaptchaReq
	4, // 1: captcha.service.v1.Captcha.GetImageCodeFromRdb:input_type -> captcha.service.v1.GetImageCodeFromRdbReq
	2, // 2: captcha.service.v1.Captcha.SendSmsCode:input_type -> captcha.service.v1.SendSmsCodeReq
	0, // 3: captcha.service.v1.Captcha.GetSmsCode:input_type -> captcha.service.v1.GetSmsCodeReq
	7, // 4: captcha.service.v1.Captcha.GetCaptcha:output_type -> captcha.service.v1.GetCaptchaReply
	5, // 5: captcha.service.v1.Captcha.GetImageCodeFromRdb:output_type -> captcha.service.v1.GetImageCodeFromRdbReply
	3, // 6: captcha.service.v1.Captcha.SendSmsCode:output_type -> captcha.service.v1.SendSmsCodeReply
	1, // 7: captcha.service.v1.Captcha.GetSmsCode:output_type -> captcha.service.v1.GetSmsCodeReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_captcha_proto_init() }
func file_v1_captcha_proto_init() {
	if File_v1_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsCodeReq); i {
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
		file_v1_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsCodeReply); i {
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
		file_v1_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeReq); i {
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
		file_v1_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeReply); i {
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
		file_v1_captcha_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageCodeFromRdbReq); i {
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
		file_v1_captcha_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageCodeFromRdbReply); i {
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
		file_v1_captcha_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaptchaReq); i {
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
		file_v1_captcha_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaptchaReply); i {
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
			RawDescriptor: file_v1_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_captcha_proto_goTypes,
		DependencyIndexes: file_v1_captcha_proto_depIdxs,
		MessageInfos:      file_v1_captcha_proto_msgTypes,
	}.Build()
	File_v1_captcha_proto = out.File
	file_v1_captcha_proto_rawDesc = nil
	file_v1_captcha_proto_goTypes = nil
	file_v1_captcha_proto_depIdxs = nil
}
