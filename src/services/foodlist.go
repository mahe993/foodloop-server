package services

import (
	"foodloop/src/database"
	"net/http"

	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

func (*FoodlistService) GetAllForUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	foodlist, err := database.GetAllForUser(userID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, foodlist)
}

func (*FoodlistService) GetFoodlist(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	userID := r.Context().Value("userID").(string)

	foodlist, err := database.GetFoodlist(userID, foodlistID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, foodlist)
}
