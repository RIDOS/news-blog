package handler

import (
	"net/http"
)

// SetupRoutes регистрирует маршруты для приложения
func SetupRoutes() {
	http.HandleFunc("/news", getNewsHandler)
	http.HandleFunc("/news/{id}", getNewsByIDHandler)
}
