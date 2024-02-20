package services

import (
	"foodloop/src/models"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

func (*FoodlistService) GetAll(w http.ResponseWriter, r *http.Request) {
	FoodlistServices := []models.Food{
		{
			FoodID:   1,
			FoodName: "hehe",
		},
	}
	render.JSON(w, r, FoodlistServices)
}

func (*FoodlistService) GetFoodlist(w http.ResponseWriter, r *http.Request) {
	FoodlistServices := []models.Foodlist{}

	id := r.Context().Value("FoodlistServiceID").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		render.Respond(w, r, err)
	}

	for _, FoodlistService := range FoodlistServices {
		if FoodlistService.FoodlistID == idInt {
			render.JSON(w, r, FoodlistService)
			return
		}
	}
	render.Respond(w, r, "FoodlistService not found")
}
