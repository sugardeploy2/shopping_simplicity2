package entities

type Category struct {
	CategoryID   int    `gorm:"primaryKey"`
	CategoryName string `gorm:"not null"`
}
