package models

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Password string
}
