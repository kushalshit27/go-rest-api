package models

import (
	"time"
)
// Post schema of the post table
type Post struct {
    ID       int64  `json:"id"`
    Title     string `json:"title"`
    Description string `json:"description"`
    Created   time.Time  `json:"created"`
}