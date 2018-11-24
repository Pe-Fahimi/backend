package models

type Location struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	ParentID *int64    `json:"parent_id,omitempty"`
	Parent   *Location `json:"parent,omitempty"`
}
