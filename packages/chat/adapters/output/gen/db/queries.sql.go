// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package queries

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const getAllConversationsByUserID = `-- name: GetAllConversationsByUserID :many
SELECT id, name, created_at, owner_id FROM conversations WHERE owner_id = $1
`

func (q *Queries) GetAllConversationsByUserID(ctx context.Context, ownerID int64) ([]Conversation, error) {
	rows, err := q.db.QueryContext(ctx, getAllConversationsByUserID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Conversation
	for rows.Next() {
		var i Conversation
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMessagesByStatus = `-- name: GetAllMessagesByStatus :many
SELECT id, conversation_id, content, created_at, sender_id, state FROM messages WHERE conversation_id = $1  AND state = ANY($2::text[])
`

type GetAllMessagesByStatusParams struct {
	ConversationID int64
	StateIn        []string
}

func (q *Queries) GetAllMessagesByStatus(ctx context.Context, arg GetAllMessagesByStatusParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getAllMessagesByStatus, arg.ConversationID, pq.Array(arg.StateIn))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.Content,
			&i.CreatedAt,
			&i.SenderID,
			&i.State,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveConversation = `-- name: SaveConversation :one
INSERT INTO conversations (owner_id, name, created_at) VALUES ($1, $2, $3) RETURNING id
`

type SaveConversationParams struct {
	OwnerID   int64
	Name      string
	CreatedAt time.Time
}

func (q *Queries) SaveConversation(ctx context.Context, arg SaveConversationParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, saveConversation, arg.OwnerID, arg.Name, arg.CreatedAt)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const saveMessage = `-- name: SaveMessage :one
INSERT INTO messages (conversation_id, sender_id, content, created_at, state) VALUES ($1, $2, $3, $4, $5) RETURNING id
`

type SaveMessageParams struct {
	ConversationID int64
	UserID         int64
	Content        string
	CreatedAt      time.Time
	State          string
}

func (q *Queries) SaveMessage(ctx context.Context, arg SaveMessageParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, saveMessage,
		arg.ConversationID,
		arg.UserID,
		arg.Content,
		arg.CreatedAt,
		arg.State,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}