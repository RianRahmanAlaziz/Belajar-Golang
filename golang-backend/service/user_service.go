package service

import (
	"golang-backend/models"
	"golang-backend/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
