package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{
		db: db,
	}
}

func (r *WishlistRepository) CreateWishlist(wishlist *entities.Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *WishlistRepository) GetWishlistByID(wishlistID int) (*entities.Wishlist, error) {
	var wishlist entities.Wishlist
	err := r.db.First(&wishlist, wishlistID).Error
	if err != nil {
		return nil, err
	}
	return &wishlist, nil
}

func (r *WishlistRepository) UpdateWishlist(wishlist *entities.Wishlist) error {
	return r.db.Save(wishlist).Error
}

func (r *WishlistRepository) DeleteWishlist(wishlist *entities.Wishlist) error {
	return r.db.Delete(wishlist).Error
}
