package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "user=foodloopfp password=foodloopmcc dbname=foodloopdb2 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
