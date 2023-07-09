package entities

import "time"

type OrderItems struct {
	ID        int     `gorm:"primaryKey"`
	ProductID int     `gorm:"not null"`
	OrderID   int     `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
