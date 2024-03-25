package models

type Foodlist struct {
	FoodlistID       int    `json:"id"`
	FoodlistName     string `json:"title"`
	FoodlistTime     string `json:"recurringTime"`
	FoodlistDay      string `json:"recurringDay"`
	FoodlistCurrIdx  int    `json:"currentFoodIdx"`
	FoodlistCategory string `json:"category"`
	FoodlistStatus   string `json:"status"`
}

type Food struct {
	FoodID       int    `json:"id"`
	FoodName     string `json:"name"`
	Descriptions string `json:"description"`
	Category     string `json:"category"`
	FoodIndex    int    `json:"index"`
}

type CreateFoodlistRequest struct {
	Query string `json:"query"`
	Title string `json:"title"`
	Day   string `json:"recurringDay"`
	Time  string `json:"recurringTime"`
}

type GetFoodlistResponse struct {
	Foodlist Foodlist `json:"foodlist"`
	Foods    []Food   `json:"foods"`
}

type UpdateFoodlistStatusRequest struct {
	Status string `json:"status"`
}
