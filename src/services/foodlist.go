package services

import (
	"foodloop/src/database"
	"foodloop/src/models"
	"net/http"

	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

func (*FoodlistService) GetAll(w http.ResponseWriter, r *http.Request) {
	Foodlist := []models.Food{
		{
			FoodName: "hehe",
		},
	}
	render.JSON(w, r, Foodlist)
}

func (*FoodlistService) GetFoodlist(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	userID := r.Context().Value("userID").(string)

	foodlist := database.GetFoodlist(userID, foodlistID)
	render.JSON(w, r, foodlist)

	// render.Respond(w, r, "FoodlistService not found")
	// return
}
