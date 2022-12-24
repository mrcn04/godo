package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrcn04/godo/pkg/db"
	"github.com/mrcn04/godo/pkg/handlers"
)

func main() {
	init := db.InitDatabase()

	defer init.DB.Close()

	h := handlers.NewHandler(init.DB)
	r := registerRoutes(h)

	log.Println("Server started on " + init.Port)
	log.Fatal(http.ListenAndServe(":"+init.Port, r))
}

func registerRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HandleHealth).Methods("GET")
	r.HandleFunc("/todos", h.HandleGetAllTodos).Methods("GET")
	r.HandleFunc("/todos", h.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", h.HandleCreateTodo).Methods("UPDATE")
	r.HandleFunc("/todos/{id}", h.HandleCreateTodo).Methods("DELETE")

	return r
}
