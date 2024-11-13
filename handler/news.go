package handler

import (
	"encoding/json"
	"github.com/RIDOS/news-blog/internal/repository"
	"net/http"
	"strconv"
	"strings"
)

func getNewsHandler(newsRepository *repository.NewsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		news, _ := newsRepository.GetAllNews()
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(news); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}

func getNewsByIDHandler(newsRepository *repository.NewsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")

		if len(pathParts) < 3 || pathParts[1] != "news" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(pathParts[2])
		if err != nil || id < 0 {
			response := map[string]string{
				"status":  "error",
				"message": "Invalid URL",
			}

			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			}
			return
		}

		paper, err := newsRepository.GetNewsByID(id)
		if err != nil {
			if err := json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": err.Error(),
			}); err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(paper); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}
