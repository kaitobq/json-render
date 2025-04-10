//go:generate go run ../../cmd/typegen/main.go -input . -output ../generated -naming snake_case

package api

import "time"

type CreateUserResponse struct {
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
