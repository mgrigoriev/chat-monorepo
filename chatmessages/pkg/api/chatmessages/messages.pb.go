// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: api/chatmessages/messages.proto

package chatmessages

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ChatMessageInfo_RECIPIENT_TYPE int32

const (
	ChatMessageInfo_UNSPECIFIED ChatMessageInfo_RECIPIENT_TYPE = 0
	ChatMessageInfo_USER        ChatMessageInfo_RECIPIENT_TYPE = 1
	ChatMessageInfo_SERVER      ChatMessageInfo_RECIPIENT_TYPE = 2
)

// Enum value maps for ChatMessageInfo_RECIPIENT_TYPE.
var (
	ChatMessageInfo_RECIPIENT_TYPE_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "USER",
		2: "SERVER",
	}
	ChatMessageInfo_RECIPIENT_TYPE_value = map[string]int32{
		"UNSPECIFIED": 0,
		"USER":        1,
		"SERVER":      2,
	}
)

func (x ChatMessageInfo_RECIPIENT_TYPE) Enum() *ChatMessageInfo_RECIPIENT_TYPE {
	p := new(ChatMessageInfo_RECIPIENT_TYPE)
	*p = x
	return p
}

func (x ChatMessageInfo_RECIPIENT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageInfo_RECIPIENT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_api_chatmessages_messages_proto_enumTypes[0].Descriptor()
}

func (ChatMessageInfo_RECIPIENT_TYPE) Type() protoreflect.EnumType {
	return &file_api_chatmessages_messages_proto_enumTypes[0]
}

func (x ChatMessageInfo_RECIPIENT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageInfo_RECIPIENT_TYPE.Descriptor instead.
func (ChatMessageInfo_RECIPIENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{0, 0}
}

// ChatMessageInfo - детали сообщения
type ChatMessageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id - id автора сообщения
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	// user_name - имя автора сообщения
	UserName string `protobuf:"bytes,2,opt,name=user_name,proto3" json:"user_name,omitempty"`
	// recipient_type - тип получателя сообщения
	RecipientType ChatMessageInfo_RECIPIENT_TYPE `protobuf:"varint,3,opt,name=recipient_type,proto3,enum=github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo_RECIPIENT_TYPE" json:"recipient_type,omitempty"`
	// recipient_id - id получателя сообщения
	RecipientId uint64 `protobuf:"varint,4,opt,name=recipient_id,proto3" json:"recipient_id,omitempty"`
	// content - содержимое сообщения
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatMessageInfo) Reset() {
	*x = ChatMessageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageInfo) ProtoMessage() {}

func (x *ChatMessageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageInfo.ProtoReflect.Descriptor instead.
func (*ChatMessageInfo) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessageInfo) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatMessageInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ChatMessageInfo) GetRecipientType() ChatMessageInfo_RECIPIENT_TYPE {
	if x != nil {
		return x.RecipientType
	}
	return ChatMessageInfo_UNSPECIFIED
}

func (x *ChatMessageInfo) GetRecipientId() uint64 {
	if x != nil {
		return x.RecipientId
	}
	return 0
}

func (x *ChatMessageInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// ChatMessage
type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id - уникальный идентификатор сообщения
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// info - детали сообщения
	Info *ChatMessageInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ChatMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatMessage) GetInfo() *ChatMessageInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveChatMessageRequest - запрос SaveChatMessage
type SaveChatMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// info - детали сообщения
	Info *ChatMessageInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SaveChatMessageRequest) Reset() {
	*x = SaveChatMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveChatMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveChatMessageRequest) ProtoMessage() {}

func (x *SaveChatMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveChatMessageRequest.ProtoReflect.Descriptor instead.
func (*SaveChatMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{2}
}

func (x *SaveChatMessageRequest) GetInfo() *ChatMessageInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveChatMessageResponse - ответ SaveChatMessage
type SaveChatMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id - уникальный идентификатор сообщения
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveChatMessageResponse) Reset() {
	*x = SaveChatMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveChatMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveChatMessageResponse) ProtoMessage() {}

func (x *SaveChatMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveChatMessageResponse.ProtoReflect.Descriptor instead.
func (*SaveChatMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{3}
}

func (x *SaveChatMessageResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ListPrivateChatMessagesRequest - запрос на получение сообщений переписки пользователя с другим пользователем
type ListPrivateChatMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id - id пользователя
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	// other_user_id - id собеседника пользователя
	OtherUserId uint64 `protobuf:"varint,2,opt,name=other_user_id,proto3" json:"other_user_id,omitempty"`
}

func (x *ListPrivateChatMessagesRequest) Reset() {
	*x = ListPrivateChatMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPrivateChatMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrivateChatMessagesRequest) ProtoMessage() {}

func (x *ListPrivateChatMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrivateChatMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListPrivateChatMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ListPrivateChatMessagesRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListPrivateChatMessagesRequest) GetOtherUserId() uint64 {
	if x != nil {
		return x.OtherUserId
	}
	return 0
}

// ListServerChatMessagesRequest - запрос на получение сообщений сервера
type ListServerChatMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// server_id - id сервера
	ServerId uint64 `protobuf:"varint,2,opt,name=server_id,proto3" json:"server_id,omitempty"`
}

func (x *ListServerChatMessagesRequest) Reset() {
	*x = ListServerChatMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServerChatMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerChatMessagesRequest) ProtoMessage() {}

func (x *ListServerChatMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerChatMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListServerChatMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListServerChatMessagesRequest) GetServerId() uint64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

// ListChatMessagesResponse - ответ ListChatMessages
type ListChatMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// chat_messages - все сообщения
	ChatMessages []*ChatMessage `protobuf:"bytes,1,rep,name=chat_messages,proto3" json:"chat_messages,omitempty"`
}

func (x *ListChatMessagesResponse) Reset() {
	*x = ListChatMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chatmessages_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatMessagesResponse) ProtoMessage() {}

func (x *ListChatMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatmessages_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListChatMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_chatmessages_messages_proto_rawDescGZIP(), []int{6}
}

func (x *ListChatMessagesResponse) GetChatMessages() []*ChatMessage {
	if x != nil {
		return x.ChatMessages
	}
	return nil
}

var File_api_chatmessages_messages_proto protoreflect.FileDescriptor

var file_api_chatmessages_messages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67,
	0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f,
	0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe4, 0x04, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0,
	0x41, 0x02, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72,
	0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x52, 0x45, 0x43, 0x49, 0x50, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x42,
	0x0b, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x0c,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41,
	0x02, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x37, 0x0a, 0x0e, 0x52, 0x45, 0x43, 0x49, 0x50, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x3a, 0xeb, 0x01, 0x92, 0x41, 0xe7,
	0x01, 0x0a, 0xe4, 0x01, 0x2a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x31, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x20, 0x2d, 0x20, 0xd0, 0xb4, 0xd0, 0xb5, 0xd1, 0x82, 0xd0,
	0xb0, 0xd0, 0xbb, 0xd0, 0xb8, 0x20, 0xd1, 0x81, 0xd0, 0xbe, 0xd0, 0xbe, 0xd0, 0xb1, 0xd1, 0x89,
	0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x8f, 0x4a, 0x5e, 0x7b, 0x22, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x31,
	0x2c, 0x20, 0x22, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x3a, 0x20, 0x32, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x20,
	0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x7d, 0xd2, 0x01, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0xd2, 0x01, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01,
	0x0e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0xd2,
	0x01, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0xd2, 0x01,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x55, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xd7,
	0x01, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x5b, 0x92, 0x41, 0x58,
	0x0a, 0x56, 0x2a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x35, 0x53, 0x61, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x2d, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1,
	0x81, 0x20, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0xd2, 0x01, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x0a, 0x17, 0x53, 0x61, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x76, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x0d, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x7f, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x3b,
	0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chatmessages_messages_proto_rawDescOnce sync.Once
	file_api_chatmessages_messages_proto_rawDescData = file_api_chatmessages_messages_proto_rawDesc
)

func file_api_chatmessages_messages_proto_rawDescGZIP() []byte {
	file_api_chatmessages_messages_proto_rawDescOnce.Do(func() {
		file_api_chatmessages_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chatmessages_messages_proto_rawDescData)
	})
	return file_api_chatmessages_messages_proto_rawDescData
}

var file_api_chatmessages_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_chatmessages_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_chatmessages_messages_proto_goTypes = []interface{}{
	(ChatMessageInfo_RECIPIENT_TYPE)(0),    // 0: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo.RECIPIENT_TYPE
	(*ChatMessageInfo)(nil),                // 1: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo
	(*ChatMessage)(nil),                    // 2: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessage
	(*SaveChatMessageRequest)(nil),         // 3: github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageRequest
	(*SaveChatMessageResponse)(nil),        // 4: github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageResponse
	(*ListPrivateChatMessagesRequest)(nil), // 5: github.com.mgrigoriev.chat_monorepo.chatmessages.ListPrivateChatMessagesRequest
	(*ListServerChatMessagesRequest)(nil),  // 6: github.com.mgrigoriev.chat_monorepo.chatmessages.ListServerChatMessagesRequest
	(*ListChatMessagesResponse)(nil),       // 7: github.com.mgrigoriev.chat_monorepo.chatmessages.ListChatMessagesResponse
}
var file_api_chatmessages_messages_proto_depIdxs = []int32{
	0, // 0: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo.recipient_type:type_name -> github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo.RECIPIENT_TYPE
	1, // 1: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessage.info:type_name -> github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo
	1, // 2: github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageRequest.info:type_name -> github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessageInfo
	2, // 3: github.com.mgrigoriev.chat_monorepo.chatmessages.ListChatMessagesResponse.chat_messages:type_name -> github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_chatmessages_messages_proto_init() }
func file_api_chatmessages_messages_proto_init() {
	if File_api_chatmessages_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chatmessages_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageInfo); i {
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
		file_api_chatmessages_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_api_chatmessages_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveChatMessageRequest); i {
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
		file_api_chatmessages_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveChatMessageResponse); i {
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
		file_api_chatmessages_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPrivateChatMessagesRequest); i {
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
		file_api_chatmessages_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServerChatMessagesRequest); i {
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
		file_api_chatmessages_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatMessagesResponse); i {
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
			RawDescriptor: file_api_chatmessages_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_chatmessages_messages_proto_goTypes,
		DependencyIndexes: file_api_chatmessages_messages_proto_depIdxs,
		EnumInfos:         file_api_chatmessages_messages_proto_enumTypes,
		MessageInfos:      file_api_chatmessages_messages_proto_msgTypes,
	}.Build()
	File_api_chatmessages_messages_proto = out.File
	file_api_chatmessages_messages_proto_rawDesc = nil
	file_api_chatmessages_messages_proto_goTypes = nil
	file_api_chatmessages_messages_proto_depIdxs = nil
}
