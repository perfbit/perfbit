// pkg/service/user_service.go
package service

import (
	"github.com/maulikam/auth-service/pkg/model"
	"github.com/maulikam/auth-service/pkg/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, nil
	}
	return user, nil
}
