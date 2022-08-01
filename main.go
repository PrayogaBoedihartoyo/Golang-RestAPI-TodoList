package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/router"
	"net/http"
)

func main() {
	router := router.Router()
	fmt.Println("Application is running...")
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
