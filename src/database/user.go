package database

import "foodloop/src/models"

func GetUser(userID string) (models.User, error) {
	rows, err := db.Query(`
    SELECT *
    FROM foodloop.people p
    WHERE p.peopleID = $1;
`, userID)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		if err := rows.Scan(
			&user.UserID,
			&user.Username,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	rows, err := db.Query(`
    SELECT *
    FROM foodloop.people p
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(
			&user.UserID,
			&user.Username,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(name string) (models.User, error) {
	rows, err := db.Query(`
	INSERT INTO foodloop.people (name)
	VALUES ($1)
	RETURNING *;
`, name)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		if err := rows.Scan(
			&user.UserID,
			&user.Username,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
