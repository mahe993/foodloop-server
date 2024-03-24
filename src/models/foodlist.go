package models

type Foodlist struct {
	FoodlistID      int    `json:"foodlistID"`
	FoodlistName    string `json:"foodlistName"`
	FoodlistTime    string `json:"foodlistTime"`
	FoodlistDay     string `json:"foodlistDay"`
	FoodlistCurrIdx int    `json:"foodlistCurrIdx"`
}

type Food struct {
	FoodID       int    `json:"foodID"`
	FoodName     string `json:"foodName"`
	Descriptions string `json:"descriptions"`
	Category     string `json:"category"`
	FoodIndex    int    `json:"foodIndex"`
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
