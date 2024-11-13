package handler

import (
	"github.com/RIDOS/news-blog/internal/repository"
	"net/http"
)

func SetupRoutes(storage *repository.NewsRepository) {
	http.HandleFunc("/news", getNewsHandler(storage))
	http.HandleFunc("/news/{id}", getNewsByIDHandler(storage))
}
