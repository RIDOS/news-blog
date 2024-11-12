package handler

import (
	"encoding/json"
	"github.com/RIDOS/news-blog/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getNewsHandler(w http.ResponseWriter, r *http.Request) {
	// test case
	news := []models.New{
		{
			Id:        1,
			Title:     "Test",
			Body:      "Test body",
			Image:     "file://",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(news); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func getNewsByIDHandler(w http.ResponseWriter, r *http.Request) {
	// test case
	news := []models.New{
		{
			Id:        1,
			Title:     "Test",
			Body:      "Test body",
			Image:     "file://",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[1] != "news" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil || id >= len(news) || id < 0 {
		response := map[string]string{
			"status":  "error",
			"message": "Invalid URL",
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	paper := news[id]
	if err := json.NewEncoder(w).Encode(paper); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
