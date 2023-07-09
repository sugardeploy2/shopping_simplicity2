package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order *entities.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetOrderByID(orderID int) (*entities.Order, error) {
	var order entities.Order
	err := r.db.Preload("OrderItems").First(&order, orderID).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateOrder(order *entities.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) DeleteOrder(order *entities.Order) error {
	return r.db.Delete(order).Error
}
