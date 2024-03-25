package database

import (
	"errors"
	"fmt"
	"foodloop/src/models"
)

func InsertFoodlist(id int, list []models.Food, time string, day string) (models.Foodlist, error) {
	if len(list) == 0 {
		return models.Foodlist{}, errors.New("empty list")
	}

	r := db.QueryRow(
		`
		INSERT INTO foodloop.foodlist(foodlistName, foodlistTime, foodlistDay, foodlistCurrIdx, foodlistCategory, foodlistStatus)
		VALUES($1, $2, $3, $4, $5)
		RETURNING *
		`,
		list[0].Category,
		time,
		day,
		0,
		list[0].Category,
		"play",
	)

	var foodlist models.Foodlist
	if err := r.Scan(
		&foodlist.FoodlistID,
		&foodlist.FoodlistName,
		&foodlist.FoodlistTime,
		&foodlist.FoodlistDay,
		&foodlist.FoodlistCurrIdx,
		&foodlist.FoodlistCategory,
		&foodlist.FoodlistStatus,
	); err != nil {
		return models.Foodlist{}, err
	}

	if _, err := db.Exec(
		`
		INSERT INTO foodloop.peopleToFoodlist(peopleID, foodlistID)
		VALUES($1, $2)
		`,
		id,
		foodlist.FoodlistID,
	); err != nil {
		return models.Foodlist{}, err
	}

	stmt := "INSERT INTO foodloop.foodlistToFood(foodlistID, foodID, foodIndex) VALUES"
	for i, v := range list {
		stmt += fmt.Sprintf("(%d, %d, %d),", foodlist.FoodlistID, v.FoodID, i+1)
	}
	stmt = stmt[0 : len(stmt)-1]
	if _, err := db.Exec(
		stmt,
	); err != nil {
		return models.Foodlist{}, err
	}

	return foodlist, nil
}

func GetAllForUser(userID string) ([]models.Foodlist, error) {
	rows, err := db.Query(
		`
		SELECT fl.*
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
			&foodlist.FoodlistCategory,
			&foodlist.FoodlistStatus,
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
		SELECT fl.*
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
			&foodlist.FoodlistCategory,
			&foodlist.FoodlistStatus,
		); err != nil {
			return models.Foodlist{}, err
		}
	}
	return foodlist, nil
}

func UpdateFoodlistStatus(id string, status string) error {
	// Construct the UPDATE query
	query := `
        UPDATE foodloop.foodlist
        SET foodlistStatus = $1
        WHERE foodlistID = $2
    `
	// Execute the update query
	_, err := db.Exec(query, status, id)
	if err != nil {
		return err
	}

	return nil
}
