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
	log.Fatal(http.ListenAndServe(":8000", router))
}
