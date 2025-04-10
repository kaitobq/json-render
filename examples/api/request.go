//go:generate go run ../../cmd/typegen/main.go -input . -output ../generated -naming snake_case

package api

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
