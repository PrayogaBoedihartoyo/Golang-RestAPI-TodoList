package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"main/config"
	"main/helper"
	"net/http"
)

type Todo struct {
	Id          int64  `json:"id_todo"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	defer db.Close()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err helper.Error
		err = helper.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = helper.GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	sqlStatement := `INSERT INTO users (email, password) VALUES ($1, $2)`
	ctx := context.Background()
	_, err = db.ExecContext(ctx, sqlStatement, user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	defer db.Close()

	var RequestUser Authentication

	err := json.NewDecoder(r.Body).Decode(&RequestUser)
	if err != nil {
		var err helper.Error
		err = helper.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var user User
	db.QueryRow("select email,password from users").Scan(&user.Email, &user.Password)

	if user.Email == "" {
		var err helper.Error
		err = helper.SetError(err, "Username or Password is incorrect/gaboleh kosong")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := helper.CheckPasswordHash(RequestUser.Password, user.Password)

	if !check {
		var err helper.Error
		err = helper.SetError(err, "Username or Password is incorrect/salah")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := helper.GenerateJWT(user.Email)
	if err != nil {
		var err helper.Error
		err = helper.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token helper.Token
	token.Email = user.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func CreateTodo(todo Todo) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO todo (status, description) VALUES ($1, $2) `

	var id int64
	ctx := context.Background()
	_, err := db.ExecContext(ctx, sqlStatement, todo.Status, todo.Description)
	if err != nil {
		log.Fatalf("cannot execute query. %v", err)
	}

	return id
}

func FindAllTodo() ([]Todo, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM todo`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("cannot execute query. %v", err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Status, &todo.Description)
		if err != nil {
			log.Fatalf("cannot get data. %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func FindTodo(id int64) (Todo, error) {
	db := config.CreateConnection()
	sqlStatement := "SELECT id, status, description FROM todo WHERE id = $1"
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	todo := Todo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Status, &todo.Description)
		if err != nil {
			panic(err)
		}
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}

}

func UpdateTodo(id int64, todo Todo) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE todo SET status=$2, description=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, todo.Status, todo.Description)

	if err != nil {
		log.Fatalf("cannot execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error Update data. %v", err)
	}

	fmt.Printf("Data updated %v\n", rowsAffected)

	return rowsAffected
}

func DeleteTodo(id int64) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM todo WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("cannot execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	//fmt.Printf("data deleted %v", rowsAffected)

	return rowsAffected
}
