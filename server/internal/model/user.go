package model

type UserEntity struct {
	Id       string
	Name     string
	Password string
}

type UserInfo struct {
	Name string
}

type UserAuth struct {
	Name     string
	Password string
}

type UserId struct {
	Id string
}
