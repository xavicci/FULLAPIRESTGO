package models

import "time"

type Post struct {
	Id          string    `json:"id"`
	PostContent string    `json:"post_content"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      string    `json:"user_id"`
}
