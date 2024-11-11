package storage

import "errors"

var (
	ErrNewsNotFound = errors.New("url not found")
	ErrNewsExists   = errors.New("news already exists")
)
