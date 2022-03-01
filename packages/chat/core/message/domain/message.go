package domain

import (
	"errors"
	"time"
)

type State int64

var messageStateMap = map[State]string{
	0: "unknown",
	1: "sent",
	2: "delivered",
	3: "read",
}

func (s State) String() string {
	return messageStateMap[s]
}

var (
	ErrMessageNotSent      = errors.New("message not sent")
	ErrMessageNotDelivered = errors.New("message not delivered")
)

const (
	Unknown State = iota
	Sent
	Delivered
	Read
)

type Message struct {
	ID             int64
	ConversationID int64
	SenderID       int64
	Content        string
	CreatedAt      time.Time
	State          State
}

func (m Message) IsUnread() bool {
	return m.State != Read
}

func (m *Message) Delivered() error {
	if m.State != Sent {
		return ErrMessageNotSent
	}

	m.State = Delivered

	return nil
}

func (m *Message) Read() error {
	if m.State != Delivered {
		return ErrMessageNotDelivered
	}

	m.State = Read

	return nil
}
