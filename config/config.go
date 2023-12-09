package config

import (
	"database/sql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
