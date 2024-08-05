-- name: CreateFeed :one
INSERT INTO
    feeds (name, url, user_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET
    last_fetched_at = now(),
    updated_at = now()
WHERE
    id = $1
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1 AND user_id = $2;