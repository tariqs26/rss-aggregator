-- name: CreateFeedFollow :one
INSERT INTO
    feed_follows (feed_id, user_id)
VALUES ($1, $2) RETURNING *;

-- name: GetFeedFollows :many
SELECT
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.feed_id,
    ff.user_id,
    f.name,
    f.url,
    f.created_at as feed_created_at,
    f.updated_at as feed_updated_at
FROM feed_follows ff
    INNER JOIN feeds f ON ff.feed_id = f.id
WHERE
    ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2;
-- Ensure the user owns the feed follow