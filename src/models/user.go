package models

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}
