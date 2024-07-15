// pkg/service/user_service.go
package service

import (
	"github.com/maulikam/auth-service/pkg/model"
	"github.com/maulikam/auth-service/pkg/repository"
	"github.com/maulikam/auth-service/pkg/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, nil
	}
	return user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}
