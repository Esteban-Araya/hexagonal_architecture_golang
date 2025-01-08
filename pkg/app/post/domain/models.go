package domain

import (
	"time"
)

type CreatePostModel struct {
	UserID    int       `json:"-" `
	Title     string    `json:"title" validate:"min=4, max=80"`
	Content   string    `json:"content"  validate:"min=10, max=1000"`
	Longitude  float64	`json:"longitude" validate:"nonzero"	`
	Latitude   float64	`json:"latitude" validate:"nonzero"`
	CreatedAt time.Time `json:"-"`
}
		
type GetPost struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"time"`
}

type PostUpdate struct {
	ID       int    `json:"id"`
	UserID   int    `json:"-"`
	Postname string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
