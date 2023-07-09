package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) CreateCategory(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetCategoryByID(categoryID int) (*entities.Category, error) {
	var category entities.Category
	err := r.db.First(&category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) DeleteCategory(category *entities.Category) error {
	return r.db.Delete(category).Error
}
