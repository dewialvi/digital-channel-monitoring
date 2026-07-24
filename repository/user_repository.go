package repository

import (
	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}