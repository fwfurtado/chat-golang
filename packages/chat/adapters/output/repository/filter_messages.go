package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	"chat/packages/chat/core/conversation/ports"
	"chat/packages/chat/core/message/domain"
	"context"
	"fmt"
)

func (r conversationRepository) FilterMessage(params ports.MessageByState) ([]*domain.Message, error) {
	ctx := context.TODO()

	statesAsString := messageStateToString(params.InState)

	filterParams := queries.GetAllMessagesByStatusParams{
		ConversationID: params.ConversationID,
		StateIn:        statesAsString,
	}

	dbMessages, err := queries.New(r.db).GetAllMessagesByStatus(ctx, filterParams)

	if err != nil {
		return nil, fmt.Errorf("error getting messages in states %v for conversation id %d from database: %w", statesAsString, params.ConversationID, err)
	}

	messages := make([]*domain.Message, len(dbMessages))

	for i, dbMessage := range dbMessages {
		messages[i] = messageToDomain(dbMessage)
	}

	return messages, nil
}
