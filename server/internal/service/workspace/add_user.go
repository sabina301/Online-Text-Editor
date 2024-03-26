package workspace

import (
	"log"
	"strconv"
)

func (s *workspaceService) AddUser(workspaceId string, userId string) error {
	wIdInt, err := strconv.Atoi(workspaceId)
	uIdInt, err := strconv.Atoi(userId)
	log.Println(wIdInt, userId)
	if err != nil {
		return err
	}
	return s.workspaceRepository.AddUser(wIdInt, uIdInt)
}
