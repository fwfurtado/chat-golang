package ports

import (
	conversationDomain "chat/packages/chat/core/conversation/domain"
	messageDomain "chat/packages/chat/core/message/domain"
)

type ConversationForUser struct {
	CurrentUserID int64
}

type MessageByState struct {
	ConversationID int64
	CurrentUserID  int64
	InState        []messageDomain.State
}

type ConversationRepository interface {
	Save(conversation *conversationDomain.Conversation) error
	ListConversations(params ConversationForUser) ([]*conversationDomain.Conversation, error)
	FilterMessage(params MessageByState) ([]*messageDomain.Message, error)
	AddMessage(message *messageDomain.Message) error
}
