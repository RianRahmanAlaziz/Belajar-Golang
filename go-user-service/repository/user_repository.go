package repository

import (
	"go-user-service/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]models.Users, error) {

	var users []models.Users

	err := r.DB.Find(&users).Error

	return users, err
}

func (r *UserRepository) Create(user *models.Users) error {

	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.Users, error) {

	var user models.Users

	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
