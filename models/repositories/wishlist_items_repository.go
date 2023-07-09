package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type WishlistItemsRepository struct {
	db *gorm.DB
}

func NewWishlistItemsRepository(db *gorm.DB) *WishlistItemsRepository {
	return &WishlistItemsRepository{
		db: db,
	}
}

func (r *WishlistItemsRepository) CreateWishlistItem(wishlistItem *entities.WishlistItem) error {
	return r.db.Create(wishlistItem).Error
}

func (r *WishlistItemsRepository) GetWishlistItemByID(wishlistItemID int) (*entities.WishlistItem, error) {
	var wishlistItem entities.WishlistItem
	err := r.db.First(&wishlistItem, wishlistItemID).Error
	if err != nil {
		return nil, err
	}
	return &wishlistItem, nil
}

func (r *WishlistItemsRepository) UpdateWishlistItem(wishlistItem *entities.WishlistItem) error {
	return r.db.Save(wishlistItem).Error
}

func (r *WishlistItemsRepository) DeleteWishlistItem(wishlistItem *entities.WishlistItem) error {
	return r.db.Delete(wishlistItem).Error
}
