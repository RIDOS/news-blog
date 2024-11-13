package repository

import (
	"errors"
	"github.com/RIDOS/news-blog/internal/models"
	"time"
)

type NewsRepository struct {
}

func NewNewsRepository() *NewsRepository {
	return &NewsRepository{}
}

func (r *NewsRepository) GetNewsByID(id int) (*models.News, error) {
	news := []models.News{
		{
			Id:        1,
			Title:     "Test",
			Body:      "Test body",
			Image:     "file://",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	if len(news) <= id {
		return nil, errors.New("id out of range")
	}

	return &news[id], nil
}

func (r *NewsRepository) GetAllNews() ([]models.News, error) {
	news := []models.News{
		{
			Id:        1,
			Title:     "Test",
			Body:      "Test body",
			Image:     "file://",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	return news, nil
}
