package models

import "time"

type Product struct {
	ID        uint
	UpdatedAt time.Time
	CreatedAt time.Time

	Name  string
	Price float64
	Image string
}
