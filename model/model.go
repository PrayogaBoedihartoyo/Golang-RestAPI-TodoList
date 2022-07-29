package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"main/config"
)

type Todo struct {
	Id          int64  `json:"id_todo"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func CreateTodo(todo Todo) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	sqlStatement := `INSERT INTO todo (status, description) VALUES ($1, $2) `

	// id yang dimasukkan akan disimpan di id ini
	var id int64
	ctx := context.Background()
	// Scan function akan menyimpan insert id didalam id id
	db.ExecContext(ctx, sqlStatement, todo.Status, todo.Description)
	//if err != nil {
	//	log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	//}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

func FindAllTodo() ([]Todo, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM todo`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Status, &todo.Description)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func FindTodo(id int64) (Todo, error) {
	db := config.CreateConnection()
	defer db.Close()

	var todo Todo
	sqlStatement := `SELECT * FROM todo WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&todo.Id, &todo.Status, &todo.Description)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return todo, nil
	case nil:
		return todo, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return todo, err
}

func UpdateTodo(id int64, todo Todo) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE todo SET status=$2, description=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, todo.Status, todo.Description)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func DeleteTodo(id int64) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM todo WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
