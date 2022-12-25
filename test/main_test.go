package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrcn04/godo/pkg/db"
	"github.com/mrcn04/godo/pkg/handlers"
	"github.com/mrcn04/godo/pkg/models"
)

func TestHealthCheck(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	handlers.HandleHealth(response, request)

	var got models.Health
	want := models.Health{Status: "OK", Date: time.Now()}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Health, '%v'", response.Body, err)
	}

	if response.Code != http.StatusOK {
		t.Errorf("Response code was %v; want 200", response.Code)
	}

	if got.Status != want.Status {
		t.Errorf("failed dude")
	}
}

func TestTodos(t *testing.T) {
	init := db.InitDatabase("../.env")

	defer init.DB.Close()

	var latestCreatedTodoId int64

	t.Run("test create todo", func(t *testing.T) {
		var body = []byte(`{"text":"testing"}`)
		request, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		handlers.NewHandler(init.DB).HandleCreateTodo(response, request)

		var got models.Todo
		err := json.NewDecoder(response.Body).Decode(&got)
		request.Body.Close()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if response.Code != http.StatusCreated {
			t.Errorf("Response code was %v; want 200", response.Code)
		}

		if got.Text != "testing" {
			t.Errorf("got %v want %v", got, string(body))
		}

		latestCreatedTodoId = got.ID
	})

	t.Run("test get all todos", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/todos", nil)
		response := httptest.NewRecorder()

		handlers.NewHandler(init.DB).HandleGetAllTodos(response, request)

		todos := make([]models.Todo, 0)
		err := json.NewDecoder(response.Body).Decode(&todos)

		if err != nil {
			t.Fatalf(response.Body.String(), err.Error())
		}

		if response.Code != http.StatusOK {
			t.Errorf("Response code was %v; want 200", response.Code)
		}

		if len(todos) <= 0 {
			t.Errorf("Got no todos")
		}
	})

	t.Run("test update todo", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/todos", nil)
		response := httptest.NewRecorder()

		q := request.URL.Query()
		q.Add("id", fmt.Sprint(latestCreatedTodoId))
		q.Add("text", "updated")
		request.URL.RawQuery = q.Encode()

		handlers.NewHandler(init.DB).HandleUpdateTodo(response, request)

		var got models.Todo
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf(response.Body.String(), err.Error())
		}

		if response.Code != http.StatusOK {
			t.Errorf("Response code was %v; want 200", response.Code)
		}

		if got.Text != "updated" {
			t.Errorf("got %v want %v", got, string(q.Get("text")))
		}
	})

	t.Run("test delete todo", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/todos/{id}", nil)
		response := httptest.NewRecorder()

		request = mux.SetURLVars(request, map[string]string{
			"id": fmt.Sprint(latestCreatedTodoId),
		})

		handlers.NewHandler(init.DB).HandleDeleteTodo(response, request)

		var got = struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: true,
			Message: "Item deleted.",
		}

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf(response.Body.String(), err.Error())
		}

		if response.Code != http.StatusOK {
			t.Errorf("Response code was %v; want 200", response.Code)
		}

		if !got.Success || got.Message != "Item deleted." {
			t.Errorf("got %v want %v", got, "")
		}
	})
}
