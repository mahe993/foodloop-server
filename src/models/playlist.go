package models

type Foodlist struct {
	FoodlistID   int    `json:"id"`
	Foodlist []Food
}

type Food struct {
	FoodID	int		`json:"foodID"`
	FoodName string `json:"foodName"`
	Descriptions string `json:"descriptions"`
}
