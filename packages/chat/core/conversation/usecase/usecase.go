package usecase

import "chat/packages/chat/core/conversation/ports"

type conversationUseCase struct {
	repository ports.ConversationRepository
}

func New(repository ports.ConversationRepository) ports.ConversationUseCase {
	return &conversationUseCase{
		repository: repository,
	}
}
