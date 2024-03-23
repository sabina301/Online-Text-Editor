package workspace

import (
	desc "Online-Text-Editor/server/pkg/workspace_v1"
	"context"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	return &desc.AddUserResponse{
		Message: "add",
	}, nil
}
