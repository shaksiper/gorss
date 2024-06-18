-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, deleted, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feed;
