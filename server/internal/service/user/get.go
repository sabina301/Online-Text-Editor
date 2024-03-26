package user

import (
	"Online-Text-Editor/server/internal/model"
	"strconv"
)

func (u *userService) Get(id string) (*model.UserInfo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepository.Get(intId)
	return user, nil
}
