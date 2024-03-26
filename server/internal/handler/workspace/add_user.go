package workspace

import (
	"Online-Text-Editor/server/internal/converter"
	desc "Online-Text-Editor/server/pkg/workspace_v1"
	"context"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	err := i.workspaceService.AddUser(req.GetWorkspaceId(), converter.ToUserId(req).Id)
	if err != nil {
		return nil, err
	}
	return &desc.AddUserResponse{
		Message: "add",
	}, nil
}
