package usecase

import (
	"chat/packages/chat/core/conversation/ports"
	messageDomain "chat/packages/chat/core/message/domain"
)

var unreadStates = []messageDomain.State{
	messageDomain.Sent,
	messageDomain.Delivered,
}

func (c conversationUseCase) PullUnread(userID, conversationID int64) ([]*ports.MessageView, error) {
	filterParam := ports.MessageByState{
		CurrentUserID:  userID,
		ConversationID: conversationID,
		InState:        unreadStates,
	}

	messages, err := c.repository.FilterMessage(filterParam)

	if err != nil {
		return nil, err
	}

	unreadMessages := make([]*ports.MessageView, len(messages))

	for i, message := range messages {
		unreadMessages[i] = messageToView(message)
	}

	return unreadMessages, nil
}
