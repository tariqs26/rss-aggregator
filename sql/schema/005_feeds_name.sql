-- +goose Up
ALTER TABLE feeds ADD COLUMN name TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE feeds DROP COLUMN name;