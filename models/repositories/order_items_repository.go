package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type OrderItemsRepository struct {
	db *gorm.DB
}

func NewOrderItemsRepository(db *gorm.DB) *OrderItemsRepository {
	return &OrderItemsRepository{
		db: db,
	}
}

func (r *OrderItemsRepository) CreateOrderItem(orderItem *entities.OrderItems) error {
	return r.db.Create(orderItem).Error
}

func (r *OrderItemsRepository) GetOrderItemByID(orderItemID int) (*entities.OrderItems, error) {
	var orderItem entities.OrderItems
	err := r.db.First(&orderItem, orderItemID).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemsRepository) UpdateOrderItem(orderItem *entities.OrderItems) error {
	return r.db.Save(orderItem).Error
}

func (r *OrderItemsRepository) DeleteOrderItem(orderItem *entities.OrderItems) error {
	return r.db.Delete(orderItem).Error
}
