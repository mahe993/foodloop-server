package database

import (
	"fmt"
	"foodloop/src/models"
)

func GetFoodlist(userID string, foodlistID string) models.Foodlist {
	rows, err := db.Query(
		`
		SELECT foodname, f.descriptions 
		FROM foodloop.people p 
		LEFT JOIN foodloop.peopleToFoodlist ptf 
		ON p.peopleid = ptf.peopleid 
		LEFT JOIN foodlistToFood ftf 
		ON ptf.foodlistid = ftf.foodlistid 
		LEFT JOIN food f 
		ON ftf.foodid = f.foodid 
		LEFT JOIN restaurantToFood rtf 
		ON rtf.foodid = f.foodid 
		LEFT JOIN restaurant r 
		ON r.restaurantid = rtf.restaurantid
		WHERE p.peopleid = $1
		AND ptf.foodlistid = $2;
		`,
		userID,
		foodlistID,
	)
	if err != nil {
		fmt.Println(err)
	}

	foodlist := []models.Food{}
	for rows.Next() {
		food := models.Food{}
		if err := rows.Scan(
			&food.FoodName,
			&food.Descriptions,
		); err != nil {
			fmt.Println(err)
		}
		foodlist = append(foodlist, food)
	}

	return models.Foodlist{
		Foodlist: foodlist,
	}

}
