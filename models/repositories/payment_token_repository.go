package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type PaymentTokenRepository struct {
	db *gorm.DB
}

func NewPaymentTokenRepository(db *gorm.DB) *PaymentTokenRepository {
	return &PaymentTokenRepository{
		db: db,
	}
}

func (r *PaymentTokenRepository) CreatePaymentToken(paymentToken *entities.PaymentToken) error {
	return r.db.Create(paymentToken).Error
}

func (r *PaymentTokenRepository) GetPaymentTokenByID(paymentTokenID int) (*entities.PaymentToken, error) {
	var paymentToken entities.PaymentToken
	err := r.db.First(&paymentToken, paymentTokenID).Error
	if err != nil {
		return nil, err
	}
	return &paymentToken, nil
}

func (r *PaymentTokenRepository) UpdatePaymentToken(paymentToken *entities.PaymentToken) error {
	return r.db.Save(paymentToken).Error
}

func (r *PaymentTokenRepository) DeletePaymentToken(paymentToken *entities.PaymentToken) error {
	return r.db.Delete(paymentToken).Error
}
