package services

import (
	"encoding/json"
	"foodloop/src/database"
	"foodloop/src/models"
	"io"
	"net/http"

	"github.com/go-chi/render"
)

type UserService struct{}

var User UserService

func (*UserService) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, users)
}

func (*UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(string)
	user, err := database.GetUser(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, user)
}

func (*UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	var req models.CreateUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	user, err := database.CreateUser(req.Name)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, user)
}
