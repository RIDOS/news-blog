package repository

import (
	"errors"
	"github.com/RIDOS/news-blog/pkg/models"
	"github.com/RIDOS/news-blog/storage"
)

type NewsRepository struct {
	st storage.Storage
}

func NewNewsRepository(st storage.Storage) *NewsRepository {
	return &NewsRepository{st: st}
}

func (r *NewsRepository) GetByID(id int) (*models.New, error) {
	news, err := r.st.GetNews(id)

	if err != nil {
		return nil, err
	}

	if news.Id <= 0 {
		return nil, storage.ErrNewsNotFound
	}

	return &news, nil
}

func (r *NewsRepository) GetAll(limit, offset int) ([]interface{}, error) {
	news, err := r.st.GetAllNews(limit, offset)
	if err != nil {
		return nil, err
	}

	// Преобразуем []models.New в []interface{}
	items := make([]interface{}, len(news))
	for i, n := range news {
		items[i] = n
	}

	return items, nil
}

func (r *NewsRepository) Count() (int, error) {
	count, err := r.st.Count()

	if err != nil {
		return 0, err
	}

	if count <= 0 {
		return 0, errors.New("count is not 0")
	}

	return count, nil
}
