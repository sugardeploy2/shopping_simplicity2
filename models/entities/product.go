package entities

import "time"

type Product struct {
	ID          int    `gorm:"primaryKey"`
	ProductName string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	CategoryID  int
	Brand       string
	Quantity    int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
