package usecase

import (
	"chat/packages/chat/core/conversation/ports"
	"fmt"
)

func (c conversationUseCase) ListConversations(userID int64) ([]*ports.ConversationSummary, error) {
	filterParams := ports.ConversationForUser{
		CurrentUserID: userID,
	}

	conversations, err := c.repository.ListConversations(filterParams)

	if err != nil {
		return nil, fmt.Errorf("failed to list conversations: %w", err)
	}

	listOfConversations := make([]*ports.ConversationSummary, len(conversations))

	for i, conversation := range conversations {
		listOfConversations[i] = conversationToSummary(conversation)
	}

	return listOfConversations, nil
}
