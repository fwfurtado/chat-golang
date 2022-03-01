package usecase

import (
	"chat/packages/chat/core/conversation/domain"
	"chat/packages/chat/core/conversation/ports"
	messageDomain "chat/packages/chat/core/message/domain"
	"time"
)

func createParamToConversation(params *ports.CreateConversationParam) *domain.Conversation {
	return &domain.Conversation{
		Name:      params.Name,
		OwnerID:   params.OwnerID,
		CreatedAt: time.Now().UTC(),
	}
}

func conversationToSummary(conversation *domain.Conversation) *ports.ConversationSummary {
	return &ports.ConversationSummary{
		ID:          conversation.ID,
		Name:        conversation.Name,
		UnreadCount: conversation.TotalUnread(),
	}
}

func messageToView(message *messageDomain.Message) *ports.MessageView {
	return &ports.MessageView{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		Content:        message.Content,
		CreatedAt:      message.CreatedAt,
		State:          message.State.String(),
	}
}

func newMessageParamToMessage(message *ports.NewMessageParam) *messageDomain.Message {
	return &messageDomain.Message{
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		Content:        message.Content,
		CreatedAt:      time.Now().UTC(),
		State:          messageDomain.Sent,
	}
}
