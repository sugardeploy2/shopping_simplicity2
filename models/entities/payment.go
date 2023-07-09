package entities

import "time"

type Payment struct {
	ID                    int       `gorm:"primaryKey"`
	UserID                int       `gorm:"not null"`
	OrderID               int       `gorm:"not null"`
	Amount                float64   `gorm:"not null"`
	PaymentDate           time.Time `gorm:"not null"`
	PaymentStatus         string    `gorm:"not null"`
	PaymentGatewayTransID string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
