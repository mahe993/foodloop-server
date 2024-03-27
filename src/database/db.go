package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db = dbConn
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
