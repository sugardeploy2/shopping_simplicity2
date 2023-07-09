package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetProductByID(productID int) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(product *entities.Product) error {
	return r.db.Delete(product).Error
}
