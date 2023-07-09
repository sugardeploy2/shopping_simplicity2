package entities

type Address struct {
	ID      int    `gorm:"primaryKey"`
	UserID  int    `gorm:"not null"`
	Address string `gorm:"not null"`
}
