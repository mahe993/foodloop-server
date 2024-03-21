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
		); err != nil {
			return nil, err
		}
		res = append(res, food)
	}
	return res, nil
}

func InsertFoodlist(id int, list []models.Food, title string, time string, day string) error {
	var newID int64
	if err := db.QueryRow(
		`
		INSERT INTO foodloop.foodlist(foodlistName, foodlistTime, foodlistDay, foodlistCurrIdx)
		VALUES($1, $2, $3, $4)
		RETURNING foodlistID
		`,
		title,
		time,
		day,
		0,
	).Scan(&newID); err != nil {
		return err
	}

	if _, err := db.Exec(
		`
		INSERT INTO foodloop.peopleToFoodlist(peopleID, foodlistID)
		VALUES($1, $2)
		`,
		id,
		newID,
	); err != nil {
		return err
	}

	stmt := "INSERT INTO foodloop.foodlistToFood(foodlistID, foodID, foodIndex) VALUES"
	for i, v := range list {
		stmt += fmt.Sprintf("(%d, %d, %d),", newID, v.FoodID, i+1)
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
		SELECT fl.foodlistID, fl.foodlistName, fl.foodlistTime, fl.foodlistDay, fl.foodlistCurrIdx
		FROM foodlist fl
		LEFT JOIN peopleToFoodlist ptfl
		ON fl.foodlistID = ptfl.foodlistID
		WHERE ptfl.peopleID = $1
		`,
		userID,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	foodlists := []models.Foodlist{}
	for rows.Next() {
		foodlist := models.Foodlist{}
		if err := rows.Scan(
			&foodlist.FoodlistID,
			&foodlist.FoodlistName,
			&foodlist.FoodlistTime,
			&foodlist.FoodlistDay,
			&foodlist.FoodlistCurrIdx,
		); err != nil {
			fmt.Println(err)
		}
		foodlists = append(foodlists, foodlist)
	}

	return foodlists, nil
}

func GetFoodlist(userID string, foodlistID string) (models.Foodlist, error) {
	rows, err := db.Query(
		`
		SELECT fl.foodlistID, fl.foodlistName, fl.foodlistTime, fl.foodlistDay, fl.foodlistCurrIdx
		FROM foodlist fl
		LEFT JOIN peopleToFoodlist ptfl
		ON fl.foodlistID = ptfl.foodlistID
		WHERE ptfl.peopleID = $1 
		AND ptfl.foodlistID = $2
		`,
		userID,
		foodlistID,
	)
	if err != nil {
		fmt.Println(err)
		return models.Foodlist{}, err
	}

	foodlist := models.Foodlist{}
	for rows.Next() {
		if err := rows.Scan(
			&foodlist.FoodlistID,
			&foodlist.FoodlistName,
			&foodlist.FoodlistTime,
			&foodlist.FoodlistDay,
			&foodlist.FoodlistCurrIdx,
		); err != nil {
			fmt.Println(err)
		}
	}

	rows, err = db.Query(
		`
		SELECT f.foodID, f.foodName, f.descriptions, ftf.foodIndex
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
		ORDER BY ftf.foodIndex ASC
		`,
		userID,
		foodlistID,
	)
	if err != nil {
		fmt.Println(err)
		return models.Foodlist{}, err
	}

	foods := []models.Food{}
	for rows.Next() {
		food := models.Food{}
		if err := rows.Scan(
			&food.FoodID,
			&food.FoodName,
			&food.Descriptions,
			&food.FoodIndex,
		); err != nil {
			fmt.Println(err)
			return models.Foodlist{}, err
		}
		foods = append(foods, food)
	}

	foodlist.Foodlist = foods
	return foodlist, nil

}
