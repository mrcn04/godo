package main

import (
	"log"
	"net/http"

	"github.com/mrcn04/godo/pkg/db"
	"github.com/mrcn04/godo/pkg/handlers"
)

func main() {
	init := db.InitDatabase()

	registerRoutes()

	log.Println("Server started on " + init.Port)
	log.Fatal(http.ListenAndServe(":"+init.Port, nil))
}

func registerRoutes() {
	http.HandleFunc("/health", handlers.HandleHealth)
}
