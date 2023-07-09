package repositories

import (
	"errors"
	"shop_khordad/models/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

/*
	func (r *UserRepository) CreateUser(user *entities.User) error {
		return r.db.Create(user).Error
	}
*/
func (ur *UserRepository) CreateUser(user *entities.User) error {
	err := ur.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID int) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *entities.User) error {
	return r.db.Save(user).Error
}

/*func (ur *UserRepository) VerifyCredential(username, password string) (*entities.User, error) {
	// Retrieve the user from the database based on the username
	var user entities.User
	err := ur.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			return nil, errors.New("Invalid credentials")
		}
		// Other database error occurred
		return nil, err
	}


	// Compare the provided password with the hashed password stored in the user entity
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Passwords do not match
		return nil, errors.New("Invalid credentials")
	}

	// Passwords match, return the user entity
	return &user, nil
}
*/

func (ur *UserRepository) VerifyCredential(email, password string) (*entities.User, error) {
	// Retrieve the user from the database based on the email
	var user entities.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			return nil, errors.New("Invalid credentials")
		}
		// Other database error occurred
		return nil, err
	}

	// Compare the password provided with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Invalid password
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}

func (r *UserRepository) DeleteUser(user *entities.User) error {
	return r.db.Delete(user).Error
}
