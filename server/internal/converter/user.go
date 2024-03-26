package converter

import (
	"Online-Text-Editor/server/internal/model"
	desc "Online-Text-Editor/server/pkg/user_v1"
)

func ToUserAuthFromDesc(info *desc.UserAuth) *model.UserAuth {
	return &model.UserAuth{
		Name:     info.Name,
		Password: info.Password,
	}
}

func ToUserInfoDesc(name string) *desc.UserInfo {
	return &desc.UserInfo{
		Name: name,
	}
}
