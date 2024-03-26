package user

import (
	"Online-Text-Editor/server/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

func (u *userService) Login(user *model.UserAuth) (string, int, error) {
	user1, err := u.userRepository.GetByUsername(user.Name, GeneratePasswordHash(user.Password))
	if err != nil {
		return "", -1, err
	}
	token, userId, err := GenerateToken(user1)
	if err != nil {
		return "", -1, err
	}
	return token, userId, nil
}

func GenerateToken(user *model.UserEntity) (string, int, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("jwtKey")))
	if err != nil {
		return "", -1, err
	}
	userIdInt, _ := strconv.Atoi(user.Id)
	return tokenStr, userIdInt, nil
}
