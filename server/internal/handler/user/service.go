package user

import (
	"Online-Text-Editor/server/internal/model"
	desc "Online-Text-Editor/server/pkg/user_v1"
)

type Service interface {
	Create(*model.UserAuth) (string, error)
	Get() model.UserInfo
	Delete() int
}

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService Service
}

func NewImplementation(userService Service) *Implementation {
	return &Implementation{userService: userService}
}
