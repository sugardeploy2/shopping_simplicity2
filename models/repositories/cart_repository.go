package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) CreateCart(cart *entities.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) GetCartByID(cartID int) (*entities.Cart, error) {
	var cart entities.Cart
	err := r.db.Preload("CartItems").First(&cart, cartID).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) UpdateCart(cart *entities.Cart) error {
	return r.db.Save(cart).Error
}

func (r *CartRepository) DeleteCart(cart *entities.Cart) error {
	return r.db.Delete(cart).Error
}
