package workspace

import (
	desc "Online-Text-Editor/server/pkg/workspace_v1"
	"context"
)

func (i *Implementation) GetWorkspace(ctx context.Context, req *desc.GetWorkspaceRequest) (*desc.GetWorkspaceResponse, error) {
	return &desc.GetWorkspaceResponse{
		Id:   "1",
		Name: "ss",
	}, nil
}
