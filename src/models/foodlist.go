package models

type Foodlist struct {
	FoodlistID      int    `json:"FoodlistID"`
	FoodlistName    string `json:"FoodlistName"`
	FoodlistTime    string `json:"FoodlistTime"`
	FoodlistDay     string `json:"FoodlistDay"`
	FoodlistCurrIdx int    `json:"FoodlistCurrIdx"`
	Foodlist        []Food
}

type Food struct {
	FoodID       int    `json:"foodID"`
	FoodName     string `json:"foodName"`
	Descriptions string `json:"descriptions"`
	FoodIndex    int    `json:"foodIndex"`
}
