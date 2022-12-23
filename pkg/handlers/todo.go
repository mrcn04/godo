package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/mrcn04/godo/pkg/models"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	todo := models.Todo{}
	log.Println(todo)
}

func (h *Handler) HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("getting all todos...")

	rows, err := h.db.Query("select * from todos")

	if err != nil {
		fmt.Printf("ERROR:: ", err.Error())
	}

	log.Println(rows)
}

func (h *Handler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {

}
