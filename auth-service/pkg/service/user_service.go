package service

import (
	"github.com/maulikam/perfbit/auth-service/pkg/model"
	"github.com/maulikam/perfbit/auth-service/pkg/repository"
	"github.com/maulikam/perfbit/auth-service/pkg/utils"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		log.Println("Error finding user:", err)
		return nil, err
	}
	if user == nil {
		log.Println("User not found")
		return nil, nil
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		log.Println("Password does not match")
		return nil, nil
	}

	if !user.Verified {
		log.Println("User not verified")
		return nil, nil
	}

	return user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) VerifyUser(username, code string) error {
	return s.Repo.VerifyUser(username, code)
}

func (s *UserService) UpdateRefreshToken(username, refreshToken string) error {
	return s.Repo.UpdateRefreshToken(username, refreshToken)
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.Repo.FindByUsername(username)
}
