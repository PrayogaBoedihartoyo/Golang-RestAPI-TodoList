package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Application is running at http://localhost:8000")
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
