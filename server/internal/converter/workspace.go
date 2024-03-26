package converter

import (
	"Online-Text-Editor/server/internal/model"
	desc "Online-Text-Editor/server/pkg/workspace_v1"
)

func ToWorkspaceWithoutId(info *desc.WorkspaceWithoutId) *model.WorkspaceWithoutId {
	return &model.WorkspaceWithoutId{
		Name: info.Name,
	}
}

func ToUserId(info *desc.AddUserRequest) *model.UserId {
	return &model.UserId{
		Id: info.UserId,
	}
}
