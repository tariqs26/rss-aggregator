-- name: CreatePost :one
INSERT INTO posts (title, description, url, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserPosts :many
SELECT posts.*
FROM posts
    JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;