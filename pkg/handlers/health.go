package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mrcn04/godo/pkg/models"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	health := models.Health{Status: "OK", Date: time.Now()}

	h, err := json.Marshal(health)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(h)
}
