package entities

type Brand struct {
	ID         int    `gorm:"primaryKey"`
	BrandName  string `gorm:"not null"`
	CategoryID int    `gorm:"not null"`
}
