package services

import (
	"errors"
	"go-user-service/models"
	"go-user-service/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUsers() ([]models.Users, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *models.Users) error {

	existingUser, _ := s.repo.FindByEmail(user.Email)

	if existingUser != nil {
		return errors.New("email already exists")
	}

	return s.repo.Create(user)
}
