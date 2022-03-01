package ports

import (
	"chat/packages/chat/core/conversation/domain"
	messageDomain "chat/packages/chat/core/message/domain"
	"time"
)

type CreateConversationParam struct {
	Name    string
	OwnerID int64
}

type NewMessageParam struct {
	ConversationID int64
	SenderID       int64
	Content        string
}

type ConversationFilter struct {
	ConversationID int64
	UserID         int64
}

type ConversationSummary struct {
	ID          int64
	Name        string
	UnreadCount int
}

type MessageView struct {
	ID             int64
	ConversationID int64
	SenderID       int64
	Content        string
	CreatedAt      time.Time
	State          string
}

type CurrentUserParam struct {
	UserID int64
}

type ConversationController interface {
	Create(params *CreateConversationParam) (*ConversationSummary, error)
	ListConversations(params *CurrentUserParam) ([]*ConversationSummary, error)
	PullUnread(filter *ConversationFilter) ([]*MessageView, error)
	AppendMessage(params *NewMessageParam) (*MessageView, error)
}

type ConversationUseCase interface {
	Create(conversation *domain.Conversation) (*ConversationSummary, error)
	ListConversations(userID int64) ([]*ConversationSummary, error)
	PullUnread(userID, conversationID int64) ([]*MessageView, error)
	AppendMessage(message *messageDomain.Message) (*MessageView, error)
}
