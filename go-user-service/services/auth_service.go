package services

import "go-user-service/utils"

func (s *UserService) Login(email string) (string, error) {

	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
