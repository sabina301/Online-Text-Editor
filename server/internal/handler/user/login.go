package user

import (
	"Online-Text-Editor/server/internal/converter"
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
	"strconv"
)

func (i *Implementation) Login(ctx context.Context, req *desc.CreateRequest) (*desc.LoginResponse, error) {
	token, userId, err := i.userService.Login(converter.ToUserAuthFromDesc(req.GetUserAuth()))
	if err != nil {
		return nil, err
	}
	return &desc.LoginResponse{
		Id:    strconv.Itoa(userId),
		Token: token,
	}, nil
}
