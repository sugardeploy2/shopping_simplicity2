package entities

import "time"

type PaymentToken struct {
	ID             int       `gorm:"primaryKey"`
	TokenID        int       `gorm:"not null"`
	UserID         int       `gorm:"not null"`
	PaymentToken   string    `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
