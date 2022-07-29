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
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
