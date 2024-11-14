package storage

import (
	"errors"
	"fmt"
	"github.com/RIDOS/news-blog/internal/storage/sqlite"
	"github.com/RIDOS/news-blog/pkg/models"
)

var (
	ErrNewsNotFound      = errors.New("news not found")
	ErrNewsAlreadyExists = errors.New("news already exists")
)

func NewStorage(dbType, connectionString string) (Storage, error) {
	switch dbType {
	case "sqlite":
		return sqlite.NewSQLiteStorage(connectionString)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

type Storage interface {
	CreateNews(title, body, image string) (int64, error)
	GetNews(id int) (models.New, error)
}
