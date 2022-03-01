package repository

import (
	queries "chat/packages/chat/adapters/output/gen/db"
	conversationDomain "chat/packages/chat/core/conversation/domain"
	"chat/packages/chat/core/conversation/ports"
	"context"
	"fmt"
)

func (r conversationRepository) ListConversations(params ports.ConversationForUser) ([]*conversationDomain.Conversation, error) {
	ctx := context.TODO()

	conversations, err := queries.New(r.db).GetAllConversationsByUserID(ctx, params.CurrentUserID)

	if err != nil {
		return nil, fmt.Errorf("error getting conversations from db: %w", err)
	}

	domainConversations := make([]*conversationDomain.Conversation, len(conversations))

	for i, conversation := range conversations {
		domainConversations[i] = conversationToDomain(conversation)
	}

	return domainConversations, nil
}
