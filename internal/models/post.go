package models

import (
	"time"
)

// Post schema of the post table
type Post struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	CreatedBy   int64     `json:"created_by"`
	Updated     time.Time `json:"updated"`
	Status      bool      `json:"status"`
	User        Users     `json:"user"`
}
