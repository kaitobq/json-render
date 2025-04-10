package models

import "time"

// Post represents a blog post
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"author_id"`
	Status    string    `json:"status"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostCreateRequest represents the request body for creating a new post
type PostCreateRequest struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required,min=10"`
	Tags    []string `json:"tags" validate:"dive,min=2,max=20"`
}

// PostUpdateRequest represents the request body for updating a post
type PostUpdateRequest struct {
	Title   *string   `json:"title,omitempty" validate:"omitempty,min=3,max=100"`
	Content *string   `json:"content,omitempty" validate:"omitempty,min=10"`
	Status  *string   `json:"status,omitempty" validate:"omitempty,oneof=draft published archived"`
	Tags    *[]string `json:"tags,omitempty" validate:"omitempty,dive,min=2,max=20"`
}

// PostResponse represents the response body for post operations
type PostResponse struct {
	Post    *Post     `json:"post"`
	Author  *User     `json:"author,omitempty"`
	Message string    `json:"message,omitempty"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time"`
}
