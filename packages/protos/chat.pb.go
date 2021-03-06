// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: packages/protos/chat.proto

package pb

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

type MessageState int32

const (
	MessageState_MESSAGE_STATE_UNKNOWN   MessageState = 0
	MessageState_MESSAGE_STATE_SENT      MessageState = 1
	MessageState_MESSAGE_STATE_DELIVERED MessageState = 2
	MessageState_MESSAGE_STATE_READ      MessageState = 3
)

// Enum value maps for MessageState.
var (
	MessageState_name = map[int32]string{
		0: "MESSAGE_STATE_UNKNOWN",
		1: "MESSAGE_STATE_SENT",
		2: "MESSAGE_STATE_DELIVERED",
		3: "MESSAGE_STATE_READ",
	}
	MessageState_value = map[string]int32{
		"MESSAGE_STATE_UNKNOWN":   0,
		"MESSAGE_STATE_SENT":      1,
		"MESSAGE_STATE_DELIVERED": 2,
		"MESSAGE_STATE_READ":      3,
	}
)

func (x MessageState) Enum() *MessageState {
	p := new(MessageState)
	*p = x
	return p
}

func (x MessageState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageState) Descriptor() protoreflect.EnumDescriptor {
	return file_packages_protos_chat_proto_enumTypes[0].Descriptor()
}

func (MessageState) Type() protoreflect.EnumType {
	return &file_packages_protos_chat_proto_enumTypes[0]
}

func (x MessageState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageState.Descriptor instead.
func (MessageState) EnumDescriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{0}
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Participant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	State     MessageState           `protobuf:"varint,3,opt,name=state,proto3,enum=chat.MessageState" json:"state,omitempty"`
	Sender    *Participant           `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetState() MessageState {
	if x != nil {
		return x.State
	}
	return MessageState_MESSAGE_STATE_UNKNOWN
}

func (x *Message) GetSender() *Participant {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Message) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type FilterByParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId int64 `protobuf:"varint,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
}

func (x *FilterByParticipant) Reset() {
	*x = FilterByParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterByParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterByParticipant) ProtoMessage() {}

func (x *FilterByParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterByParticipant.ProtoReflect.Descriptor instead.
func (*FilterByParticipant) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{2}
}

func (x *FilterByParticipant) GetParticipantId() int64 {
	if x != nil {
		return x.ParticipantId
	}
	return 0
}

type ConversationSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UnreadCount int64                  `protobuf:"varint,5,opt,name=unread_count,json=unreadCount,proto3" json:"unread_count,omitempty"`
	IsNew       bool                   `protobuf:"varint,6,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	LastMessage *Message               `protobuf:"bytes,7,opt,name=last_message,json=lastMessage,proto3,oneof" json:"last_message,omitempty"`
}

func (x *ConversationSummary) Reset() {
	*x = ConversationSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationSummary) ProtoMessage() {}

func (x *ConversationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationSummary.ProtoReflect.Descriptor instead.
func (*ConversationSummary) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ConversationSummary) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConversationSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConversationSummary) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConversationSummary) GetUnreadCount() int64 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

func (x *ConversationSummary) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *ConversationSummary) GetLastMessage() *Message {
	if x != nil {
		return x.LastMessage
	}
	return nil
}

type ListConversationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conversations []*ConversationSummary `protobuf:"bytes,1,rep,name=conversations,proto3" json:"conversations,omitempty"`
}

func (x *ListConversationsResponse) Reset() {
	*x = ListConversationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConversationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConversationsResponse) ProtoMessage() {}

func (x *ListConversationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConversationsResponse.ProtoReflect.Descriptor instead.
func (*ListConversationsResponse) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ListConversationsResponse) GetConversations() []*ConversationSummary {
	if x != nil {
		return x.Conversations
	}
	return nil
}

type CreateConversationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnerId int64  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *CreateConversationRequest) Reset() {
	*x = CreateConversationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConversationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConversationRequest) ProtoMessage() {}

func (x *CreateConversationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConversationRequest.ProtoReflect.Descriptor instead.
func (*CreateConversationRequest) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{5}
}

func (x *CreateConversationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateConversationRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type FilterConversationByOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId        int64 `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ConversationId int64 `protobuf:"varint,2,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
}

func (x *FilterConversationByOwner) Reset() {
	*x = FilterConversationByOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConversationByOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConversationByOwner) ProtoMessage() {}

func (x *FilterConversationByOwner) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConversationByOwner.ProtoReflect.Descriptor instead.
func (*FilterConversationByOwner) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{6}
}

func (x *FilterConversationByOwner) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *FilterConversationByOwner) GetConversationId() int64 {
	if x != nil {
		return x.ConversationId
	}
	return 0
}

type FilterMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId int64 `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ParticipantId  int64 `protobuf:"varint,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
}

func (x *FilterMessages) Reset() {
	*x = FilterMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterMessages) ProtoMessage() {}

func (x *FilterMessages) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterMessages.ProtoReflect.Descriptor instead.
func (*FilterMessages) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{7}
}

func (x *FilterMessages) GetConversationId() int64 {
	if x != nil {
		return x.ConversationId
	}
	return 0
}

func (x *FilterMessages) GetParticipantId() int64 {
	if x != nil {
		return x.ParticipantId
	}
	return 0
}

type ListMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ListMessagesResponse) Reset() {
	*x = ListMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_protos_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResponse) ProtoMessage() {}

func (x *ListMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packages_protos_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return file_packages_protos_chat_proto_rawDescGZIP(), []int{8}
}

func (x *ListMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_packages_protos_chat_proto protoreflect.FileDescriptor

var file_packages_protos_chat_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x3c, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x79, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0xf6, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e,
	0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77,
	0x12, 0x35, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x5f, 0x0a, 0x19, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2a, 0x76, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c,
	0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x03,
	0x32, 0xcc, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x79, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_packages_protos_chat_proto_rawDescOnce sync.Once
	file_packages_protos_chat_proto_rawDescData = file_packages_protos_chat_proto_rawDesc
)

func file_packages_protos_chat_proto_rawDescGZIP() []byte {
	file_packages_protos_chat_proto_rawDescOnce.Do(func() {
		file_packages_protos_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_packages_protos_chat_proto_rawDescData)
	})
	return file_packages_protos_chat_proto_rawDescData
}

var file_packages_protos_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_packages_protos_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_packages_protos_chat_proto_goTypes = []interface{}{
	(MessageState)(0),                 // 0: chat.MessageState
	(*Participant)(nil),               // 1: chat.Participant
	(*Message)(nil),                   // 2: chat.Message
	(*FilterByParticipant)(nil),       // 3: chat.FilterByParticipant
	(*ConversationSummary)(nil),       // 4: chat.ConversationSummary
	(*ListConversationsResponse)(nil), // 5: chat.ListConversationsResponse
	(*CreateConversationRequest)(nil), // 6: chat.CreateConversationRequest
	(*FilterConversationByOwner)(nil), // 7: chat.FilterConversationByOwner
	(*FilterMessages)(nil),            // 8: chat.FilterMessages
	(*ListMessagesResponse)(nil),      // 9: chat.ListMessagesResponse
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_packages_protos_chat_proto_depIdxs = []int32{
	0,  // 0: chat.Message.state:type_name -> chat.MessageState
	1,  // 1: chat.Message.sender:type_name -> chat.Participant
	10, // 2: chat.Message.timestamp:type_name -> google.protobuf.Timestamp
	10, // 3: chat.ConversationSummary.created_at:type_name -> google.protobuf.Timestamp
	2,  // 4: chat.ConversationSummary.last_message:type_name -> chat.Message
	4,  // 5: chat.ListConversationsResponse.conversations:type_name -> chat.ConversationSummary
	2,  // 6: chat.ListMessagesResponse.messages:type_name -> chat.Message
	3,  // 7: chat.ConversationService.ListConversations:input_type -> chat.FilterByParticipant
	6,  // 8: chat.ConversationService.CreateConversation:input_type -> chat.CreateConversationRequest
	7,  // 9: chat.ConversationService.PullUnreadMessages:input_type -> chat.FilterConversationByOwner
	8,  // 10: chat.ConversationService.ListenMessages:input_type -> chat.FilterMessages
	5,  // 11: chat.ConversationService.ListConversations:output_type -> chat.ListConversationsResponse
	4,  // 12: chat.ConversationService.CreateConversation:output_type -> chat.ConversationSummary
	9,  // 13: chat.ConversationService.PullUnreadMessages:output_type -> chat.ListMessagesResponse
	2,  // 14: chat.ConversationService.ListenMessages:output_type -> chat.Message
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_packages_protos_chat_proto_init() }
func file_packages_protos_chat_proto_init() {
	if File_packages_protos_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packages_protos_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
		file_packages_protos_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_packages_protos_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterByParticipant); i {
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
		file_packages_protos_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationSummary); i {
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
		file_packages_protos_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConversationsResponse); i {
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
		file_packages_protos_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConversationRequest); i {
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
		file_packages_protos_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConversationByOwner); i {
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
		file_packages_protos_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterMessages); i {
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
		file_packages_protos_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesResponse); i {
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
	file_packages_protos_chat_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packages_protos_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packages_protos_chat_proto_goTypes,
		DependencyIndexes: file_packages_protos_chat_proto_depIdxs,
		EnumInfos:         file_packages_protos_chat_proto_enumTypes,
		MessageInfos:      file_packages_protos_chat_proto_msgTypes,
	}.Build()
	File_packages_protos_chat_proto = out.File
	file_packages_protos_chat_proto_rawDesc = nil
	file_packages_protos_chat_proto_goTypes = nil
	file_packages_protos_chat_proto_depIdxs = nil
}
