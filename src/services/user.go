package services

import (
	"foodloop/src/database"
	"net/http"

	"github.com/go-chi/render"
)

type UserService struct{}

var User UserService

func (*UserService) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		render.Respond(w, r, err)
	}

	render.JSON(w, r, users)
}

func (*UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(string)
	user, err := database.GetUser(id)
	if err != nil {
		render.Respond(w, r, err)
	}

	render.JSON(w, r, user)
}
