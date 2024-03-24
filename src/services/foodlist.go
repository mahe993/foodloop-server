package services

import (
	"encoding/json"
	"errors"
	"foodloop/src/database"
	"foodloop/src/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/bbalet/stopwords"
	"github.com/go-chi/render"
)

type FoodlistService struct{}

var Foodlist FoodlistService

func cleanQuery(s string) []string {
	cleanContent := stopwords.CleanString(s, "en", true)
	tags := strings.Split(cleanContent, " ")
	if strings.Contains(s, "chicken rice") {
		tags = append(tags, "chicken rice")
	}
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
	var req models.CreateFoodlistRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	tags := cleanQuery(req.Query)
	foods, err := database.GenerateFoodlist(tags)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	if len(foods) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "we couldn't generate a list for that query :(")
		return
	}
	fl, err := database.InsertFoodlist(userID, foods, req.Title, req.Time, req.Day)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, fl)
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
	var (
		joinedErr error
		foodlist  models.Foodlist
		foods     []models.Food
		wg        sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()

		fl, err := database.GetFoodlist(userID, foodlistID)
		if err != nil {
			errors.Join(joinedErr, err)
			return
		}
		foodlist = fl
	}()

	go func() {
		defer wg.Done()

		f, err := database.GetFoods(foodlistID)
		if err != nil {
			errors.Join(joinedErr, err)
			return
		}
		foods = f
	}()

	wg.Wait()
	if joinedErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, joinedErr.Error())
		return
	}

	res := models.GetFoodlistResponse{
		Foodlist: foodlist,
		Foods:    foods,
	}

	render.JSON(w, r, res)
}
