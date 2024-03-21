package user

import (
	"Online-Text-Editor/server/internal/model"
)

type UserService interface {
	Create(*model.UserAuth) (string, error)
	Get() model.UserInfo
	Delete() int
}

type Repository interface {
	Create(name string, passwordHash string) (int, error)
	Get()
	Delete()
}

type userService struct {
	userRepository Repository
}

func NewUserService(repo Repository) *userService {
	return &userService{
		userRepository: repo,
	}
}
