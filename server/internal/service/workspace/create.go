package workspace

import (
	"Online-Text-Editor/server/internal/model"
	"strconv"
)

func (s *workspaceService) Create(w *model.WorkspaceWithoutId, userId int) (string, error) {
	id, err := s.workspaceRepository.Create(w.Name, userId)
	return strconv.Itoa(id), err
}
