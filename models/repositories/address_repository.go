package repositories

import (
	"shop_khordad/models/entities"

	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		DB: db,
	}
}

func (r *AddressRepository) CreateAddress(address *entities.Address) error {
	return r.DB.Create(address).Error
}

func (r *AddressRepository) GetAddressByID(id uint) (*entities.Address, error) {
	address := &entities.Address{}
	if err := r.DB.First(address, id).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *AddressRepository) UpdateAddress(address *entities.Address) error {
	return r.DB.Save(address).Error
}

func (r *AddressRepository) DeleteAddress(address *entities.Address) error {
	return r.DB.Delete(address).Error
}
