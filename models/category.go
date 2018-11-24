package models

type Category struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	ParentID *int64    `json:"parent_id,omitempty"`
	Parent   *Category `json:"parent,omitempty"`
}
