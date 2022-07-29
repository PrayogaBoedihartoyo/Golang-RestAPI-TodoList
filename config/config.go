package config

import (
	"database/sql"
	"fmt"
)

// kita buat koneksi dgn db posgres
func CreateConnection() *sql.DB {

	// Kita buka koneksi ke db
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=Todo sslmode=disable password=postgres")

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to Database!")
	// return the connection
	return db
}
