package models

import "time"

type Session struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	User      *User      `json:"user,omitempty"`
	Token     string     `json:"token"`
	ClientIP  *string    `json:"client_ip,omitempty"`
	UserAgent *string    `json:"user_agent,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt time.Time  `json:"expires_at"`
	DeletedAt *time.Time `json:"-"`
}
