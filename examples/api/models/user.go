package models

import "time"

// User represents a user in the system
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCreateRequest represents the request body for creating a new user
type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	Username *string `json:"username,omitempty" validate:"omitempty,min=3,max=20"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Role     *string `json:"role,omitempty" validate:"omitempty,oneof=admin user"`
}

// UserResponse represents the response body for user operations
type UserResponse struct {
	User    *User     `json:"user"`
	Message string    `json:"message,omitempty"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time"`
}
