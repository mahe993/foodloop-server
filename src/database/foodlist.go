package database

import (
	"fmt"
	"foodloop/src/models"
	"strings"
)

func GetTagsID(tags []string) []models.Food {
	concatTags := strings.Join(tags, "','")
	rows, err := db.Query(
		`
		SELECT fff.foodName, fff.descriptions
		FROM
			(
			SELECT ff.foodName, ff.descriptions, COUNT(ff.tagName) as count
			FROM (
				SELECT f.foodName, f.descriptions, t.tagName 
				FROM foodloop.food f 
				LEFT JOIN foodloop.foodToTag ftt 
				ON f.foodID = ftt.foodID 
				LEFT JOIN foodloop.tag t 
				ON t.tagID = ftt.tagID 
				WHERE t.tagName 
				IN ('`+concatTags+`')
				) as ff
			GROUP BY ff.foodName, ff.descriptions
			) as fff
		WHERE fff.count = $1
		`,
		2,
	)
	if err != nil {
		fmt.Println(err)
	}
	res := []models.Food{}
	for rows.Next() {
		var food models.Food
		if err := rows.Scan(
			&food.FoodName,
			&food.Descriptions,
		); err != nil {
			fmt.Println(err)
		}
		res = append(res, food)
	}
	return res
}

func GetAllForUser(userID string) ([]models.Foodlist, error) {
	rows, err := db.Query(
		`
		SELECT ptf.foodlistid, foodname, f.descriptions 
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
		`,
		userID,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	foodlists := map[string]*models.Foodlist{}
	for rows.Next() {
		var id string
		food := models.Food{}
		if err := rows.Scan(
			&id,
			&food.FoodName,
			&food.Descriptions,
		); err != nil {
			fmt.Println(err)
		}
		currListStruct, ok := foodlists[id]
		if !ok {
			currListStruct = &models.Foodlist{}
			foodlists[id] = currListStruct
		}
		currListStruct.Foodlist = append(currListStruct.Foodlist, food)
	}

	response := []models.Foodlist{}
	for _, v := range foodlists {
		response = append(response, *v)
	}
	return response, nil
}

func GetFoodlist(userID string, foodlistID string) (models.Foodlist, error) {
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
		return models.Foodlist{}, err
	}

	foodlist := []models.Food{}
	for rows.Next() {
		food := models.Food{}
		if err := rows.Scan(
			&food.FoodName,
			&food.Descriptions,
		); err != nil {
			fmt.Println(err)
			return models.Foodlist{}, err
		}
		foodlist = append(foodlist, food)
	}

	return models.Foodlist{
		Foodlist: foodlist,
	}, nil

}
