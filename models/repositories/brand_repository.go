package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{
		db: db,
	}
}

func (r *BrandRepository) CreateBrand(brand *entities.Brand) error {
	return r.db.Create(brand).Error
}

func (r *BrandRepository) GetBrandByID(brandID int) (*entities.Brand, error) {
	var brand entities.Brand
	err := r.db.First(&brand, brandID).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r *BrandRepository) UpdateBrand(brand *entities.Brand) error {
	return r.db.Save(brand).Error
}

func (r *BrandRepository) DeleteBrand(brand *entities.Brand) error {
	return r.db.Delete(brand).Error
}
