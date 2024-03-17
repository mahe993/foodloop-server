package services

import (
	"encoding/json"
	"foodloop/src/database"
	"io"
	"net/http"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

type postResp struct {
	Query string `json:"query"`
}

func cleanQuery(s string) []string {
	cleanContent := stopwords.CleanString(s, "en", true)
	tags := strings.Split(cleanContent, " ")
	return tags[:len(tags)-1]
}

func (*FoodlistService) CreateFoodlist(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	var resp postResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	tags := cleanQuery(resp.Query)
	foodlist := database.GetTagsID(tags)

	render.JSON(w, r, foodlist)
	// render.Respond(w, r, "FoodlistService not found")
	// return
}

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
