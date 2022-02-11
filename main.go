package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Health struct {
	Status string
	Date   time.Time
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/health", health)
	fmt.Println("Server started on " + port)
	http.ListenAndServe(":"+port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func health(w http.ResponseWriter, r *http.Request) {
	health := Health{"OK", time.Now()}

	js, err := json.Marshal(health)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
