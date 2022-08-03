package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/helper"
	"main/model"
	"net/http"
	"strconv"
)

type response struct {
	Id          int64  `json:"id_todo,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

type Response struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Data        interface{} `json:"todo"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("Error decode data object.  %v", err)
	}
	insertID := model.CreateTodo(todo)

	res := response{
		Id:      insertID,
		Message: "Data has been created",
	}

	json.NewEncoder(w).Encode(res)
}

func FindTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("cannot convert to int.  %v", err)
	}
	todo, err := model.FindTodo(int64(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := response{
			Status:      "Data gak ada",
			Description: "Mau cari apa ?",
			Message:     "NOT FOUND",
		}
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(webResponse)
		return
	}
	res := Response{
		Status:      "Succesfully",
		Message:     "You Got it",
		Description: "Nice Work",
		Data:        todo,
	}

	json.NewEncoder(w).Encode(res)
}

func FindAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := model.FindAllTodo()

	if err != nil {
		log.Fatalf("cannot get data. %v", err)
	}

	var response Response
	response.Status = "Successfully"
	response.Message = "Success"
	response.Data = todos

	json.NewEncoder(w).Encode(response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("cannot convert string to int.  %v", err)
	}

	var todo model.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("cannot decode object.  %v", err)
	}

	_, err = model.FindTodo(int64(id))
	if err != nil {
		var err helper.Error
		err = helper.SetError(err, "There is no data record")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	updatedRows := model.UpdateTodo(int64(id), todo)
	msg := fmt.Sprintf("Todo updated. Todo %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("cannot convert string to int.  %v", err)
	}

	_, err = model.FindTodo(int64(id))
	if err != nil {
		var err helper.Error
		err = helper.SetError(err, "There is no data record")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	deletedRows := model.DeleteTodo(int64(id))

	msg := fmt.Sprintf("Todo deleted %v", deletedRows)
	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
