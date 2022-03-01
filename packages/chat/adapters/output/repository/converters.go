package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	"chat/packages/chat/core/conversation/domain"
)

func conversationToDomain(conversation queries.Conversation) *domain.Conversation {
	return &domain.Conversation{
		ID:        conversation.ID,
		Name:      conversation.Name,
		CreatedAt: conversation.CreatedAt,
		OwnerID:   conversation.OwnerID,
	}
}
