package repository

import (
	"github.com/RIDOS/news-blog/internal/storage"
	"github.com/RIDOS/news-blog/pkg/models"
	"time"
)

type NewsRepository struct {
	st storage.Storage
}

func NewNewsRepository(st storage.Storage) *NewsRepository {
	return &NewsRepository{st: st}
}

func (r *NewsRepository) GetNewsByID(id int) (*models.New, error) {
	news, err := r.st.GetNews(id)

	if err != nil {
		return nil, err
	}

	if news.Id <= 0 {
		return nil, storage.ErrNewsNotFound
	}

	return &news, nil
}

func (r *NewsRepository) GetAllNews() ([]models.New, error) {
	news := []models.New{
		{
			Id:        1,
			Title:     "Test",
			Body:      "Test body",
			Image:     "file://",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return news, nil
}
