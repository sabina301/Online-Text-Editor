package user

import "Online-Text-Editor/server/internal/model"

type Repository interface {
	Create(name string, passwordHash string) (int, error)
	Get(id int) (*model.UserInfo, error)
	GetByUsername(username string, password_hash string) (*model.UserEntity, error)
}

type userService struct {
	userRepository Repository
}

func NewUserService(repo Repository) *userService {
	return &userService{
		userRepository: repo,
	}
}
