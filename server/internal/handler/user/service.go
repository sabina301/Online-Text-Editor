package user

import (
	"Online-Text-Editor/server/internal/service/user"
	desc "Online-Text-Editor/server/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService user.UserService
}

func NewImplementation(userService user.UserService) *Implementation {
	return &Implementation{userService: userService}
}
