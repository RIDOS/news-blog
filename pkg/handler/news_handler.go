package handler

import (
	"encoding/json"
	pg "github.com/RIDOS/news-blog/internal/pagination"
	"github.com/RIDOS/news-blog/pkg/repository"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type NewsHandler struct {
	NewsRepository *repository.NewsRepository
	Logger         *slog.Logger
}

func NewNewsHandler(log *slog.Logger, newsRepo *repository.NewsRepository) *NewsHandler {
	return &NewsHandler{
		NewsRepository: newsRepo,
		Logger:         log,
	}
}

func (h *NewsHandler) SetupRoutes() {
	http.HandleFunc("/news", h.getAllNewsHandler())
	http.HandleFunc("/news/", h.getNewsByIdHandler())
}

func (h *NewsHandler) getAllNewsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination := pg.NewPagination(h.NewsRepository, 0, 100)
		news, err := pagination.CurrentPage()

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			h.Logger.Error("Failed to retrieve news", err.Error())
			http.Error(w, "Failed to retrieve news", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(news); err != nil {
			h.Logger.Error("Failed to encode JSON", err.Error())
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}

func (h *NewsHandler) getNewsByIdHandler() http.HandlerFunc {
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

		paper, err := h.NewsRepository.GetByID(id)
		if err != nil {
			h.Logger.Error("Failed to retrieve news", err.Error())
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
			h.Logger.Error("Failed to encode JSON", err.Error())
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}
