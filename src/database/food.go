package database

import (
	"foodloop/src/models"
)

func GetFoods(foodlistID string) ([]models.Food, error) {
	rows, err := db.Query(`
    SELECT f.foodID, f.foodName, f.descriptions, fftf.foodIndex
    FROM foodloop.food f
    JOIN foodloop.foodlistToFood fftf ON f.foodID = fftf.foodID
    WHERE fftf.foodlistID = $1;
`, foodlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := []models.Food{}
	for rows.Next() {
		food := models.Food{}
		if err := rows.Scan(
			&food.FoodID,
			&food.FoodName,
			&food.Descriptions,
			&food.FoodIndex,
		); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return foods, nil
}
