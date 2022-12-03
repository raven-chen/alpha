package models

import "time"

type Post struct {
	ID        uint
	UpdatedAt time.Time
	CreatedAt time.Time

	Title string
	Body  string
}
