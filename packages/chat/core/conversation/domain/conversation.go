package domain

import (
	"chat/packages/chat/core/message/domain"
	"time"
)

type Conversation struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	OwnerID   int64
	Messages  []domain.Message
}

func (c *Conversation) AddTextMessage(sentMessage domain.Message) {
	c.Messages = append(c.Messages, sentMessage)
}

func (c Conversation) TotalUnread() int {
	total := 0

	for _, message := range c.Messages {
		if message.IsUnread() {
			total++
		}
	}

	return total
}
