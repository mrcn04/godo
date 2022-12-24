package api

import (
	"github.com/gorilla/mux"
	"github.com/mrcn04/godo/pkg/handlers"
)

func RegisterRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HandleHealth).Methods("GET")
	r.HandleFunc("/todos", h.HandleGetAllTodos).Methods("GET")
	r.HandleFunc("/todos", h.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todos", h.HandleUpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", h.HandleDeleteTodo).Methods("DELETE")

	return r
}
