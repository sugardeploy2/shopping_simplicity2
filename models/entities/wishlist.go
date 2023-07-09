package entities

import "time"

type Wishlist struct {
	ID           int       `gorm:"primaryKey"`
	UserID       int       `gorm:"not null"`
	CreationDate time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
