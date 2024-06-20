-- name: CreatePost :one
INSERT INTO post (
    id, created_at, updated_at, title, description, published_at, url, feed_id
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 )
RETURNING *;

-- name: GetPostsForUser :many
SELECT post.* FROM post JOIN feed_follow ON post.feed_id = feed_follow.feed_id
WHERE feed_follow.user_id = $1
ORDER BY post.published_at DESC
LIMIT $2;
