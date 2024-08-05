-- +goose Up
CREATE TABLE feed_follows (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    feed_id INT NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;