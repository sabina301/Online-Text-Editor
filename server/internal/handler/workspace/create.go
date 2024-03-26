package workspace

import (
	"Online-Text-Editor/server/internal/converter"
	desc "Online-Text-Editor/server/pkg/workspace_v1"
	"context"
)

func (i *Implementation) CreateWorkspace(ctx context.Context, req *desc.CreateWorkspaceRequest) (*desc.CreateWorkspaceResponse, error) {
	userId := ctx.Value("userId")
	workspaceId, err := i.workspaceService.Create(converter.ToWorkspaceWithoutId(req.WorkspaceWithoutId), userId.(int))
	if err != nil {
		return nil, err
	}
	return &desc.CreateWorkspaceResponse{
		Id: workspaceId,
	}, err
}
