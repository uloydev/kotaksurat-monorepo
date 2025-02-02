// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.21.12
// source: email_handler/mailbox.proto

package email_handler_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMailboxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMailboxRequest) Reset() {
	*x = GetMailboxRequest{}
	mi := &file_email_handler_mailbox_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMailboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailboxRequest) ProtoMessage() {}

func (x *GetMailboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailboxRequest.ProtoReflect.Descriptor instead.
func (*GetMailboxRequest) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{0}
}

func (x *GetMailboxRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMailboxCredentialsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMailboxCredentialsRequest) Reset() {
	*x = GetMailboxCredentialsRequest{}
	mi := &file_email_handler_mailbox_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMailboxCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailboxCredentialsRequest) ProtoMessage() {}

func (x *GetMailboxCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailboxCredentialsRequest.ProtoReflect.Descriptor instead.
func (*GetMailboxCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{1}
}

func (x *GetMailboxCredentialsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateMailboxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMailboxRequest) Reset() {
	*x = CreateMailboxRequest{}
	mi := &file_email_handler_mailbox_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMailboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMailboxRequest) ProtoMessage() {}

func (x *CreateMailboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMailboxRequest.ProtoReflect.Descriptor instead.
func (*CreateMailboxRequest) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMailboxRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateMailboxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateMailboxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMailboxRequest) Reset() {
	*x = UpdateMailboxRequest{}
	mi := &file_email_handler_mailbox_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMailboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailboxRequest) ProtoMessage() {}

func (x *UpdateMailboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailboxRequest.ProtoReflect.Descriptor instead.
func (*UpdateMailboxRequest) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMailboxRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMailboxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteMailboxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMailboxRequest) Reset() {
	*x = DeleteMailboxRequest{}
	mi := &file_email_handler_mailbox_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMailboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMailboxRequest) ProtoMessage() {}

func (x *DeleteMailboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMailboxRequest.ProtoReflect.Descriptor instead.
func (*DeleteMailboxRequest) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMailboxRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MailboxData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	UnreadCount   int64                  `protobuf:"varint,4,opt,name=unreadCount,proto3" json:"unreadCount,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updatedAt,proto3,oneof" json:"updatedAt,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deletedAt,proto3,oneof" json:"deletedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailboxData) Reset() {
	*x = MailboxData{}
	mi := &file_email_handler_mailbox_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailboxData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxData) ProtoMessage() {}

func (x *MailboxData) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxData.ProtoReflect.Descriptor instead.
func (*MailboxData) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{5}
}

func (x *MailboxData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailboxData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MailboxData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MailboxData) GetUnreadCount() int64 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

func (x *MailboxData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MailboxData) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MailboxData) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MailboxCredentialData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ApiKey        string                 `protobuf:"bytes,5,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	SMTPUserName  string                 `protobuf:"bytes,6,opt,name=SMTPUserName,proto3" json:"SMTPUserName,omitempty"`
	SMTPPassword  string                 `protobuf:"bytes,7,opt,name=SMTPPassword,proto3" json:"SMTPPassword,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailboxCredentialData) Reset() {
	*x = MailboxCredentialData{}
	mi := &file_email_handler_mailbox_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailboxCredentialData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxCredentialData) ProtoMessage() {}

func (x *MailboxCredentialData) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxCredentialData.ProtoReflect.Descriptor instead.
func (*MailboxCredentialData) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{6}
}

func (x *MailboxCredentialData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailboxCredentialData) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *MailboxCredentialData) GetSMTPUserName() string {
	if x != nil {
		return x.SMTPUserName
	}
	return ""
}

func (x *MailboxCredentialData) GetSMTPPassword() string {
	if x != nil {
		return x.SMTPPassword
	}
	return ""
}

type MailboxResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *MailboxData           `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailboxResponse) Reset() {
	*x = MailboxResponse{}
	mi := &file_email_handler_mailbox_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxResponse) ProtoMessage() {}

func (x *MailboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxResponse.ProtoReflect.Descriptor instead.
func (*MailboxResponse) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{7}
}

func (x *MailboxResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MailboxResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MailboxResponse) GetData() *MailboxData {
	if x != nil {
		return x.Data
	}
	return nil
}

type MailboxCredentialResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *MailboxCredentialData `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailboxCredentialResponse) Reset() {
	*x = MailboxCredentialResponse{}
	mi := &file_email_handler_mailbox_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailboxCredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxCredentialResponse) ProtoMessage() {}

func (x *MailboxCredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_handler_mailbox_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxCredentialResponse.ProtoReflect.Descriptor instead.
func (*MailboxCredentialResponse) Descriptor() ([]byte, []int) {
	return file_email_handler_mailbox_proto_rawDescGZIP(), []int{8}
}

func (x *MailboxCredentialResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MailboxCredentialResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MailboxCredentialResponse) GetData() *MailboxCredentialData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_email_handler_mailbox_proto protoreflect.FileDescriptor

var file_email_handler_mailbox_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f,
	0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c,
	0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c,
	0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbf, 0x02, 0x0a, 0x0b, 0x4d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x72,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x3d, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x87, 0x01, 0x0a,
	0x15, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x4d, 0x54, 0x50, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x4d, 0x54, 0x50, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x4d, 0x54, 0x50, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x4d, 0x54, 0x50, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a,
	0x19, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd5, 0x03, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x12, 0x20, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2b, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x23, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12,
	0x23, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x23, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x6b, 0x6f, 0x74, 0x61, 0x6b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x69, 0x64, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_handler_mailbox_proto_rawDescOnce sync.Once
	file_email_handler_mailbox_proto_rawDescData = file_email_handler_mailbox_proto_rawDesc
)

func file_email_handler_mailbox_proto_rawDescGZIP() []byte {
	file_email_handler_mailbox_proto_rawDescOnce.Do(func() {
		file_email_handler_mailbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_handler_mailbox_proto_rawDescData)
	})
	return file_email_handler_mailbox_proto_rawDescData
}

var file_email_handler_mailbox_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_email_handler_mailbox_proto_goTypes = []any{
	(*GetMailboxRequest)(nil),            // 0: email_handler.GetMailboxRequest
	(*GetMailboxCredentialsRequest)(nil), // 1: email_handler.GetMailboxCredentialsRequest
	(*CreateMailboxRequest)(nil),         // 2: email_handler.CreateMailboxRequest
	(*UpdateMailboxRequest)(nil),         // 3: email_handler.UpdateMailboxRequest
	(*DeleteMailboxRequest)(nil),         // 4: email_handler.DeleteMailboxRequest
	(*MailboxData)(nil),                  // 5: email_handler.MailboxData
	(*MailboxCredentialData)(nil),        // 6: email_handler.MailboxCredentialData
	(*MailboxResponse)(nil),              // 7: email_handler.MailboxResponse
	(*MailboxCredentialResponse)(nil),    // 8: email_handler.MailboxCredentialResponse
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
}
var file_email_handler_mailbox_proto_depIdxs = []int32{
	9,  // 0: email_handler.MailboxData.createdAt:type_name -> google.protobuf.Timestamp
	9,  // 1: email_handler.MailboxData.updatedAt:type_name -> google.protobuf.Timestamp
	9,  // 2: email_handler.MailboxData.deletedAt:type_name -> google.protobuf.Timestamp
	5,  // 3: email_handler.MailboxResponse.data:type_name -> email_handler.MailboxData
	6,  // 4: email_handler.MailboxCredentialResponse.data:type_name -> email_handler.MailboxCredentialData
	0,  // 5: email_handler.Mailbox.GetMailbox:input_type -> email_handler.GetMailboxRequest
	1,  // 6: email_handler.Mailbox.GetMailboxCredentials:input_type -> email_handler.GetMailboxCredentialsRequest
	2,  // 7: email_handler.Mailbox.CreateMailbox:input_type -> email_handler.CreateMailboxRequest
	3,  // 8: email_handler.Mailbox.UpdateMailbox:input_type -> email_handler.UpdateMailboxRequest
	4,  // 9: email_handler.Mailbox.DeleteMailbox:input_type -> email_handler.DeleteMailboxRequest
	7,  // 10: email_handler.Mailbox.GetMailbox:output_type -> email_handler.MailboxResponse
	8,  // 11: email_handler.Mailbox.GetMailboxCredentials:output_type -> email_handler.MailboxCredentialResponse
	7,  // 12: email_handler.Mailbox.CreateMailbox:output_type -> email_handler.MailboxResponse
	7,  // 13: email_handler.Mailbox.UpdateMailbox:output_type -> email_handler.MailboxResponse
	7,  // 14: email_handler.Mailbox.DeleteMailbox:output_type -> email_handler.MailboxResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_email_handler_mailbox_proto_init() }
func file_email_handler_mailbox_proto_init() {
	if File_email_handler_mailbox_proto != nil {
		return
	}
	file_email_handler_mailbox_proto_msgTypes[5].OneofWrappers = []any{}
	file_email_handler_mailbox_proto_msgTypes[7].OneofWrappers = []any{}
	file_email_handler_mailbox_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_email_handler_mailbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_handler_mailbox_proto_goTypes,
		DependencyIndexes: file_email_handler_mailbox_proto_depIdxs,
		MessageInfos:      file_email_handler_mailbox_proto_msgTypes,
	}.Build()
	File_email_handler_mailbox_proto = out.File
	file_email_handler_mailbox_proto_rawDesc = nil
	file_email_handler_mailbox_proto_goTypes = nil
	file_email_handler_mailbox_proto_depIdxs = nil
}
