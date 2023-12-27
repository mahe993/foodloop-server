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
		{Id: 1, Username: "user1"},
		{Id: 2, Username: "user2"},
		{Id: 3, Username: "user3"},
	}
	render.JSON(w, r, users)
}

func (*UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{Id: 1, Username: "user1"},
		{Id: 2, Username: "user2"},
		{Id: 3, Username: "user3"},
	}

	id := r.Context().Value("userID").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		render.Respond(w, r, err)
	}

	for _, user := range users {
		if user.Id == idInt {
			render.JSON(w, r, user)
			return
		}
	}
	render.Respond(w, r, "User not found")
}
