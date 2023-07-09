package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type CartItemsRepository struct {
	db *gorm.DB
}

func NewCartItemsRepository(db *gorm.DB) *CartItemsRepository {
	return &CartItemsRepository{
		db: db,
	}
}

func (r *CartItemsRepository) CreateCartItem(cartItem *entities.CartItems) error {
	return r.db.Create(cartItem).Error
}

func (r *CartItemsRepository) GetCartItemByID(cartItemID int) (*entities.CartItems, error) {
	var cartItem entities.CartItems
	err := r.db.First(&cartItem, cartItemID).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (r *CartItemsRepository) UpdateCartItem(cartItem *entities.CartItems) error {
	return r.db.Save(cartItem).Error
}

func (r *CartItemsRepository) DeleteCartItem(cartItem *entities.CartItems) error {
	return r.db.Delete(cartItem).Error
}
