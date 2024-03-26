package user

import (
	"Online-Text-Editor/server/internal/model"
	desc "Online-Text-Editor/server/pkg/user_v1"
)

type Service interface {
	Create(*model.UserAuth) (string, error)
	Get(id string) (*model.UserInfo, error)
	Login(*model.UserAuth) (string, int, error)
}

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService Service
}

func NewImplementation(service Service) *Implementation {
	return &Implementation{
		userService: service,
	}
}
