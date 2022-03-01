
-- name: GetAllConversationsByUserID :many
SELECT * FROM conversations WHERE owner_id = @owner_id;

-- name: SaveConversation :one
INSERT INTO conversations (owner_id, name, created_at) VALUES (@owner_id, @name, @created_at) RETURNING id;

-- name: SaveMessage :one
INSERT INTO messages (conversation_id, sender_id, content, created_at, state) VALUES (@conversation_id, @user_id, @content, @created_at, @state) RETURNING id;

-- name: GetAllMessagesByStatus :many
SELECT * FROM messages WHERE conversation_id = $1  AND state = ANY(@state_in::text[]);