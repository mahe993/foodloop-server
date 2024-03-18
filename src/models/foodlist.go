package models

type Foodlist struct {
	Foodlist []Food
}

type Food struct {
	FoodID       int    `json:"foodID"`
	FoodName     string `json:"foodName"`
	Descriptions string `json:"descriptions"`
}
