package entities

import "time"

type Review struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	ProductID  int       `gorm:"not null"`
	Rating     int       `gorm:"not null"`
	ReviewText string    `gorm:"not null"`
	ReviewDate time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
