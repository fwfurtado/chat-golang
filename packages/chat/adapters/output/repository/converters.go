package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	"chat/packages/chat/core/conversation/domain"
	messageDomain "chat/packages/chat/core/message/domain"
)

func conversationToDomain(conversation queries.Conversation) *domain.Conversation {
	return &domain.Conversation{
		ID:        conversation.ID,
		Name:      conversation.Name,
		CreatedAt: conversation.CreatedAt,
		OwnerID:   conversation.OwnerID,
	}
}

func messageStateToString(states []messageDomain.State) []string {
	result := make([]string, len(states))
	for i, state := range states {
		result[i] = state.String()
	}
	return result
}

func messageToDomain(message queries.Message) *messageDomain.Message {
	return &messageDomain.Message{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		Content:        message.Content,
		CreatedAt:      message.CreatedAt,
		State:          messageDomain.StateFromString(message.State),
	}
}
