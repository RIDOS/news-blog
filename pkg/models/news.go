package models

import "time"

type New struct {
	Id        int
	Title     string
	Body      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
