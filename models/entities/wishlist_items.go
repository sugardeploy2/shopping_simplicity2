package entities

type WishlistItem struct {
	ID         int `gorm:"primaryKey"`
	WishlistID int `gorm:"not null"`
	ProductID  int `gorm:"not null"`
}
