-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;