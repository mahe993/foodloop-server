package models

type Foodlist struct {
	FoodlistID      int    `json:"foodlistID"`
	FoodlistName    string `json:"foodlistName"`
	FoodlistTime    string `json:"foodlistTime"`
	FoodlistDay     string `json:"foodlistDay"`
	FoodlistCurrIdx int    `json:"foodlistCurrIdx"`
	Foodlist        []Food
}

type Food struct {
	FoodID       int    `json:"foodID"`
	FoodName     string `json:"foodName"`
	Descriptions string `json:"descriptions"`
	FoodIndex    int    `json:"foodIndex"`
}
