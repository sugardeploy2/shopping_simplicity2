package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) CreateReview(review *entities.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) GetReviewByID(reviewID int) (*entities.Review, error) {
	var review entities.Review
	err := r.db.First(&review, reviewID).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) UpdateReview(review *entities.Review) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) DeleteReview(review *entities.Review) error {
	return r.db.Delete(review).Error
}
