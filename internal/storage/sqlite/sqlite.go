package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/RIDOS/news-blog/pkg/models"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewSQLiteStorage(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS news(
		id INTEGER PRIMARY KEY,
		title VARCHAR(50) NOT NULL UNIQUE,
		body TEXT NOT NULL,
		image VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS news_idx ON news(title);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateNews(title, body, image string) (int64, error) {
	const op = "storage.sqlite.CreateNews"

	stmt, err := s.db.Prepare("INSERT INTO news(title, body, image) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(title, body, image)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %s", op, "news already exists")
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetNews(id int) (models.New, error) {
	const op = "storage.sqlite.GetNews"

	stmt, err := s.db.Prepare("SELECT * FROM news WHERE id=?")
	if err != nil {
		return models.New{}, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Query(id)
	if err != nil {
		return models.New{}, fmt.Errorf("%s: %s", op, "news not found")
	}

	var news models.New
	if res.Next() {
		err = res.Scan(&news.Id, &news.Title, &news.Body, &news.Image, &news.CreatedAt, &news.UpdatedAt)
		if err != nil {
			return models.New{}, fmt.Errorf("%s: %w", op, err)
		}
		return news, nil
	}

	return news, nil
}
