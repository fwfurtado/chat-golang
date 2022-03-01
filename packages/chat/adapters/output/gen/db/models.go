// Code generated by sqlc. DO NOT EDIT.

package queries

import (
	"time"
)

type Conversation struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	OwnerID   int64
}

type Message struct {
	ID             int64
	ConversationID int64
	Content        string
	CreatedAt      time.Time
	SenderID       int64
	State          string
}

type User struct {
	ID   int64
	Name string
}
