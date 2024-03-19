package services

import (
	"encoding/json"
	"foodloop/src/database"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

type postResp struct {
	Query string `json:"query"`
	Title string `json:"title"`
	Day   string `json:"recurringDay"`
	Time  string `json:"recurringTime"`
}

func cleanQuery(s string) []string {
	cleanContent := stopwords.CleanString(s, "en", true)
	tags := strings.Split(cleanContent, " ")
	return tags[:len(tags)-1]
}

func (*FoodlistService) CreateFoodlist(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := strconv.Atoi(userIDStr)
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
	foodlist, err := database.GenerateFoodlist(tags)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	if err := database.InsertFoodlist(userID, foodlist, resp.Day, resp.Time, resp.Title); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, foodlist)

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
