package workspace

import (
	"Online-Text-Editor/server/internal/service/workspace"
)

type Implementation struct {
	workspaceService workspace.WorkspaceService
}

func NewImplementation(workspaceService workspace.WorkspaceService) *Implementation {
	return &Implementation{workspaceService: workspaceService}
}
