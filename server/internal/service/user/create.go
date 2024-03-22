package user

import (
	"Online-Text-Editor/server/internal/model"
	"crypto/sha256"
	"fmt"
)

const salt = "lolkekcheburek"

func (u *userService) Create(user *model.UserAuth) (string, error) {
	hashPassword := generateHashPassword(user.Password)
	id, err := u.userRepository.Create(user.Name, hashPassword)
	if err != nil {
		return "-1", err
	}
	return string(rune(id)), nil
}

func generateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
