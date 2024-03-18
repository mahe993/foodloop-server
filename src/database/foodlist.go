package database

import (
	"fmt"
	"foodloop/src/models"
	"strings"
)

func GenerateFoodlist(tags []string) ([]models.Food, error) {
	concatTags := strings.Join(tags, "','")
	rows, err := db.Query(
		`
		SELECT fff.foodID, fff.foodName, fff.descriptions
		FROM
			(
			SELECT ff.foodID, ff.foodName, ff.descriptions, COUNT(ff.tagName) as count
			FROM (
				SELECT f.foodID, f.foodName, f.descriptions, t.tagName 
				FROM foodloop.food f 
				LEFT JOIN foodloop.foodToTag ftt 
				ON f.foodID = ftt.foodID 
				LEFT JOIN foodloop.tag t 
				ON t.tagID = ftt.tagID 
				WHERE t.tagName 
				IN ('`+concatTags+`')
				) as ff
			GROUP BY ff.foodID, ff.foodName, ff.descriptions
			) as fff
		WHERE fff.count = $1
		`,
		2,
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
		); err != nil {
			return nil, err
		}
		res = append(res, food)
	}
	return res, nil
}

func InsertFoodlist(id int, list []models.Food) error {
	var currID int
	if err := db.QueryRow(
		`
		SELECT MAX(foodlistID)
		FROM foodloop.peopleToFoodlist
		`,
	).Scan(&currID); err != nil {
		return err
	}

	if _, err := db.Exec(
		`
		INSERT INTO foodloop.foodlist(foodlistID)
		VALUES($1)
		`,
		currID+1,
	); err != nil {
		return err
	}

	if _, err := db.Exec(
		`
		INSERT INTO foodloop.peopleToFoodlist(peopleID, foodlistID)
		VALUES($1, $2)
		`,
		id,
		currID+1,
	); err != nil {
		return err
	}

	stmt := "INSERT INTO foodloop.foodlistToFood(foodlistID, foodID) VALUES"
	for _, v := range list {
		stmt += fmt.Sprintf("(%d, %d),", currID+1, v.FoodID)
	}
	stmt = stmt[0 : len(stmt)-1]
	if _, err := db.Exec(
		stmt,
	); err != nil {
		return err
	}

	return nil
}

func GetAllForUser(userID string) ([]models.Foodlist, error) {
	rows, err := db.Query(
		`
		SELECT ptf.foodlistid, f.foodID, f.foodName, f.descriptions 
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
			&food.FoodID,
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
		SELECT f.foodID, f.foodName, f.descriptions 
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
			&food.FoodID,
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
