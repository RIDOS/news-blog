package handler

import (
	"encoding/json"
	"github.com/RIDOS/news-blog/internal/repository"
	"net/http"
	"strconv"
	"strings"
)

type NewsHandler struct {
	NewsRepository *repository.NewsRepository
}

func NewNewsHandler(newsRepo *repository.NewsRepository) *NewsHandler {
	return &NewsHandler{NewsRepository: newsRepo}
}

func (h *NewsHandler) SetupRoutes() {
	http.HandleFunc("/news", h.getNewsHandler())
	http.HandleFunc("/news/", h.getNewsByIDHandler())
}

func (h *NewsHandler) getNewsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		news, err := h.NewsRepository.GetAllNews()

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			http.Error(w, "Failed to retrieve news", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(news); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}

func (h *NewsHandler) getNewsByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/news/")
		id, err := strconv.Atoi(idStr)

		w.Header().Set("Content-Type", "application/json")

		if err != nil || id < 0 {
			if err := json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Invalid news ID",
			}); err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			}
			return
		}

		paper, err := h.NewsRepository.GetNewsByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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
