package config

import (
	"database/sql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=7L2HKhn.gDa=+gu host=db.fbehvmvhrduojiyrrwhi.supabase.co port=5432 dbname=postgres")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
