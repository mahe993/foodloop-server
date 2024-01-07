package services

import (
	"foodloop/src/models"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type UserService struct{}

var User UserService

func (*UserService) GetAll(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{UserID: 1, Username: "user1", Password: "123"},
		{UserID: 2, Username: "user2", Password: "123"},
		{UserID: 3, Username: "user3", Password: "123"},
	}
	render.JSON(w, r, users)
}

func (*UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{UserID: 1, Username: "user1"},
		{UserID: 2, Username: "user2"},
		{UserID: 3, Username: "user3"},
	}

	id := r.Context().Value("userID").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		render.Respond(w, r, err)
	}

	for _, user := range users {
		if user.UserID == idInt {
			render.JSON(w, r, user)
			return
		}
	}
	render.Respond(w, r, "User not found")
}
