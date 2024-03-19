package user

import (
	"Online-Text-Editor/server/internal/model"
	"Online-Text-Editor/server/internal/repository/user"
)

type UserService interface {
	Create(*model.UserAuth) (string, error)
	Get() model.UserInfo
	Delete() int
}

type userService struct {
	userRepository user.UserRepository
}

func NewUserService(repo user.UserRepository) *userService {
	return &userService{
		userRepository: repo,
	}
}
