package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	"chat/packages/chat/core/conversation/domain"
	"context"
	"fmt"
)

func (r conversationRepository) Save(conversation *domain.Conversation) error {
	ctx := context.TODO()
	conversationParams := queries.SaveConversationParams{
		Name:      conversation.Name,
		CreatedAt: conversation.CreatedAt,
		OwnerID:   conversation.OwnerID,
	}

	savedConversationID, err := queries.New(r.db).SaveConversation(ctx, conversationParams)

	if err != nil {
		return fmt.Errorf("error saving conversation to database: %w", err)
	}

	conversation.ID = savedConversationID

	return nil
}
