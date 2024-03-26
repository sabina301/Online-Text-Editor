package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCurrentUserId(ctx *gin.Context) (int, error) {
	userId, err := ctx.Get("userId")
	if err != true {
		return -1, errors.New("cant get userId")
	}
	userIdInt := userId.(int)
	return userIdInt, nil
}

func SetCurrentUserId(ctx *gin.Context, userId int) error {
	ctx.Set("userId", userId)
	return nil
}

// parsing token
func UserIdentity(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Cant get metadata")
	}
	method := strings.Split(info.FullMethod, "/")[2]
	authHeader, ok := md["authorization"]
	if method == "CreateWorkspace" || method == "AddUser" {
		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		userId, err := ParseToken(tokenString)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, "userId", userId)
	}
	return handler(ctx, req)
}

func ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("jwtKey")), nil
	})

	if err != nil {
		log.Println("Invalid token: ", err)
		return -1, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)
		id, err := strconv.Atoi(userId)
		if err != nil {
			return -1, err
		}
		return id, nil
	} else {
		return -1, err
	}

}
