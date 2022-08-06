package config

import (
	"database/sql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=Todo sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
