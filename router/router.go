package router

import (
	"github.com/gorilla/mux"
	"main/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/todo", controller.FindAllTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.FindTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", controller.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.DeleteTodo).Methods("DELETE", "OPTIONS")

	return router
}
