package models

import "time"

type Session struct {
	ID        int64
	UserID    int64
	User      *User
	Token     string
	ClientIP  *string
	UserAgent *string
	CreatedAt time.Time
	ExpiresAt time.Time
	DeletedAt *time.Time
}
