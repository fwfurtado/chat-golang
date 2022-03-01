package repository

import (
	"chat/packages/chat/core/conversation/ports"
	"database/sql"
)

type conversationRepository struct {
	db *sql.DB
}

func New(db *sql.DB) ports.ConversationRepository {
	return &conversationRepository{db}
}
