package user

import (
	"Online-Text-Editor/server/internal/converter"
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(converter.ToUserAuthFromDesc(req.GetUserAuth()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
