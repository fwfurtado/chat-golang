package usecase

import (
	"chat/packages/chat/core/conversation/domain"
	"chat/packages/chat/core/conversation/ports"
	"fmt"
)

func (c conversationUseCase) Create(conversation *domain.Conversation) (*ports.ConversationSummary, error) {

	err := c.repository.Save(conversation)

	if err != nil {
		return nil, fmt.Errorf("failed to create conversation: %w", err)
	}

	return conversationToSummary(conversation), nil
}
