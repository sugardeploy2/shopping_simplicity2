package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) CreatePayment(payment *entities.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) GetPaymentByID(paymentID int) (*entities.Payment, error) {
	var payment entities.Payment
	err := r.db.First(&payment, paymentID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) UpdatePayment(payment *entities.Payment) error {
	return r.db.Save(payment).Error
}

func (r *PaymentRepository) DeletePayment(payment *entities.Payment) error {
	return r.db.Delete(payment).Error
}
