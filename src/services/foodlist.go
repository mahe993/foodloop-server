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
	var cleanedTags []string
	if strings.Contains(s, "chicken rice") {
		cleanedTags = append(cleanedTags, "chicken rice")
		s = strings.Replace(s, "chicken rice", "", -1)
	}
	cleanContent := stopwords.CleanString(s, "en", true)
	tags := strings.Split(cleanContent, " ")

	cat := map[string]string{
		"burgers":  "burger",
		"pastas":   "pasta",
		"biryanis": "biryani",
		"laksas":   "laksa",
	}
	for i, tag := range tags {
		if val, ok := cat[tag]; ok {
			tags[i] = val
		}
	}
	cleanedTags = append(cleanedTags, tags...)

	return cleanedTags[:len(cleanedTags)-1]
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
	foods, err := database.GenerateFoods(tags)
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
	fl, err := database.InsertFoodlist(userID, foods, req.Time, req.Day)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, fl)
}

func (*FoodlistService) GetAllForUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	foodlists, err := database.GetAllForUser(userID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, foodlists)
}

func (*FoodlistService) GetFoodlist(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	userID := r.Context().Value("userID").(string)
	var (
		foodlist              models.Foodlist
		foods                 []models.Food
		foodlistErr, foodsErr error
		wg                    sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()

		fl, err := database.GetFoodlist(userID, foodlistID)
		if err != nil {
			foodlistErr = err
			return
		}
		foodlist = fl
	}()

	go func() {
		defer wg.Done()

		f, err := database.GetFoods(foodlistID)
		if err != nil {
			foodsErr = err
			return
		}
		foods = f
	}()

	wg.Wait()
	err := errors.Join(foodlistErr, foodsErr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	res := models.GetFoodlistResponse{
		Foodlist: foodlist,
		Foods:    foods,
	}

	render.JSON(w, r, res)
}

func (*FoodlistService) UpdateFoodlistStatus(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	var req models.UpdateFoodlistStatusRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	f, err := database.UpdateFoodlistStatus(foodlistID, req.Status)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, f)
}

func (*FoodlistService) UpdateFoodlistIndex(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	var req models.UpdateFoodlistIndexRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	f, err := database.UpdateFoodlistIndex(foodlistID, req.Index)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, f)
}

func (*FoodlistService) DeleteFoodlist(w http.ResponseWriter, r *http.Request) {
	foodlistID := r.Context().Value("foodlistID").(string)
	userID := r.Context().Value("userID").(string)

	err := database.DeleteFoodlist(foodlistID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	foodlists, err := database.GetAllForUser(userID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, foodlists)

}
