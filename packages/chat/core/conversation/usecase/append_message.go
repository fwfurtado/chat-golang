package usecase

import (
	"chat/packages/chat/core/conversation/ports"
	"chat/packages/chat/core/message/domain"
	"fmt"
)

func (c conversationUseCase) AppendMessage(message *domain.Message) (*ports.MessageView, error) {

	err := c.repository.AddMessage(message)

	if err != nil {
		return nil, fmt.Errorf("failed to append message: %v", err)
	}

	return messageToView(message), nil
}
