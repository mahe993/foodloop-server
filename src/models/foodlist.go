package models

type Foodlist struct {
	Foodlist []Food
}

type Food struct {
	FoodName     string `json:"foodName"`
	Descriptions string `json:"descriptions"`
}
