package user

import (
	"Online-Text-Editor/server/internal/converter"
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(req.GetId())

	if err != nil {
		return nil, err
	}

	userD := converter.ToUserInfoDesc(user.Name)
	return &desc.GetResponse{
		UserInfo: userD,
	}, nil
}
