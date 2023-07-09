package entities

import "time"

type Order struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	OrderDate   time.Time `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
