package repository

import (
	"golang-backend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	err := r.DB.Find(&users).Error

	return users, err

}
func (r *UserRepository) Create(user *models.User) error {

	return r.DB.Create(user).Error
}
