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

	foodlist := database.GetAllForUser(userID)
	render.JSON(w, r, foodlist)
}

func (*FoodlistService) GetFoodlist(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	userID := r.Context().Value("userID").(string)

	foodlist := database.GetFoodlist(userID, foodlistID)
	render.JSON(w, r, foodlist)

	// render.Respond(w, r, "FoodlistService not found")
	// return
}
