package models

import "time"

const (
	ItemStatusPending   = "pending"
	ItemStatusPublished = "published"
	ItemStatusRejected  = "rejected"
)

type Item struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	AuthorID   int64      `json:"author_id"`
	Author     User       `json:"author,omitempty"`
	LocationID int64      `json:"location_id"`
	Location   Location   `json:"location,omitempty"`
	CategoryID int64      `json:"category_id"`
	Category   Category   `json:"category,omitempty"`
	ImageURL   *string    `json:"image_url,omitempty"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"-"`
}
