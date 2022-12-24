package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrcn04/godo/pkg/models"
	"github.com/mrcn04/godo/pkg/utils"
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
	var t models.Todo

	id := r.FormValue("id")
	text := r.FormValue("text")

	if id == "" || text == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.db.QueryRow(
		"UPDATE todos set text = $1, updated_at = NOW() WHERE id = $2 returning id, updated_at, created_at",
		text, id,
	).Scan(&t.ID, &t.Updated, &t.Created)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	t.Text = text
	utils.RespondWithJSON(w, http.StatusCreated, t)
}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := h.db.Exec("DELETE from todos WHERE id = $1", id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var message string
	if count > 0 {
		message = "Item deleted."
	} else {
		message = "No item has been deleted."
	}

	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: message,
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}
