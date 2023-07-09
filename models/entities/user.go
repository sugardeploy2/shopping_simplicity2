package entities

import "time"

type User struct {
	ID                int    `gorm:"primaryKey"`
	FirstName         string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	Email             string `gorm:"unique;not null"`
	Password          string `gorm:"not null"`
	PhoneNumber       string
	Gender            string
	DateOfBirth       time.Time
	RegistrationDate  time.Time
	LastLoginDate     time.Time
	AccountStatus     string
	PreferredLanguage string
	Status            string
	PaymentMethod     string
	PaymentID         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
