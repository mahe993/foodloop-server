package database

import (
	"foodloop/src/models"
	"strings"
)

func GenerateFoods(tags []string) ([]models.Food, error) {
	concatTags := strings.Join(tags, "','")
	rows, err := db.Query(
		`
		SELECT fff.foodID, fff.foodName, fff.descriptions, fff.category
		FROM
			(
			SELECT ff.foodID, ff.foodName, ff.descriptions, ff.category, COUNT(ff.tagName) as count
			FROM (
				SELECT f.foodID, f.foodName, f.descriptions, t.tagName, f.category
				FROM foodloop.food f 
				LEFT JOIN foodloop.foodToTag ftt 
				ON f.foodID = ftt.foodID 
				LEFT JOIN foodloop.tag t 
				ON t.tagID = ftt.tagID 
				WHERE t.tagName 
				IN ('`+concatTags+`')
				) as ff
			GROUP BY ff.foodID, ff.foodName, ff.descriptions, ff.category
			) as fff
		WHERE fff.count = $1
		`,
		1,
	)
	if err != nil {
		return nil, err
	}
	res := []models.Food{}
	for rows.Next() {
		var food models.Food
		if err := rows.Scan(
			&food.FoodID,
			&food.FoodName,
			&food.Descriptions,
			&food.Category,
		); err != nil {
			return nil, err
		}
		res = append(res, food)
	}
	return res, nil
}

func GetFoods(foodlistID string) ([]models.Food, error) {
	rows, err := db.Query(`
    SELECT f.foodID, f.foodName, f.descriptions, f.category, fftf.foodIndex
    FROM foodloop.food f
    JOIN foodloop.foodlistToFood fftf ON f.foodID = fftf.foodID
    WHERE fftf.foodlistID = $1
	ORDER BY fftf.foodIndex;
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
			&food.Category,
			&food.FoodIndex,
		); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return foods, nil
}
