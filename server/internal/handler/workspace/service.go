package workspace

import (
	"Online-Text-Editor/server/internal/model"
	desc "Online-Text-Editor/server/pkg/workspace_v1"
)

type Service interface {
	Create(ww *model.WorkspaceWithoutId, userId int) (string, error)
	Get() (int, error)
	AddUser(workspaceId string, userId string) error
}

type Implementation struct {
	desc.UnimplementedWorkspaceV1Server
	workspaceService Service
}

func NewImplementation(workspaceService Service) *Implementation {
	return &Implementation{workspaceService: workspaceService}
}
