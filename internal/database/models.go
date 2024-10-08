// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            int32
	Url           string
	UserID        uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    int32
	UserID    uuid.UUID
}

type Post struct {
	ID          int32
	Title       string
	Description sql.NullString
	Url         string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FeedID      int32
}

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ApiKey    string
}
