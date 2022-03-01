package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	"chat/packages/chat/core/message/domain"
	"context"
	"fmt"
)

func (r conversationRepository) AddMessage(message *domain.Message) error {

	ctx := context.TODO()

	messageParams := queries.SaveMessageParams{
		ConversationID: message.ConversationID,
		UserID:         message.SenderID,
		CreatedAt:      message.CreatedAt,
		Content:        message.Content,
		State:          message.State.String(),
	}

	saveMessageID, err := queries.New(r.db).SaveMessage(ctx, messageParams)

	if err != nil {
		return fmt.Errorf("failed to save message to database: %w", err)
	}

	message.ID = saveMessageID

	return nil
}
