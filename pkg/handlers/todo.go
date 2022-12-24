package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mrcn04/godo/pkg/models"
	"github.com/mrcn04/godo/utils"
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
	var t models.Todo

	timestamp := time.Now()
	t.Text = r.FormValue("text")
	t.Created = timestamp.String()
	t.Updated = timestamp.String()

	if r.FormValue("text") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.db.QueryRow(
		"INSERT INTO todos(text, created_at, updated_at) VALUES($1, $2, $3) returning id",
		t.Text, timestamp, timestamp,
	).Scan(&t.ID)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, t)
}

func (h *Handler) HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("select * from todos")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo

		err := rows.Scan(&t.ID, &t.Text, &t.Created, &t.Updated)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		todos = append(todos, t)
	}

	utils.RespondWithJSON(w, http.StatusOK, todos)
}

func (h *Handler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {

}
