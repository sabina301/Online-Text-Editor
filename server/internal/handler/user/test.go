package user

import (
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
	"log"
)

func (i *Implementation) Test(ctx context.Context, req *desc.TestRequest) (*desc.TestResponse, error) {
	log.Println("HI!")
	return &desc.TestResponse{Str: "Hello!"}, nil
}
