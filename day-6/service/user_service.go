package service

import (
	"fmt"
)

type UserService interface {
	GetListUsers() string
	CreateUser() string
	GetUserById(userId int) string
	UpdateUserById(userId int, payload string) string
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) GetListUsers() string {
	return "Get list users works!"
}

func (u *userService) CreateUser() string {
	return "Create new user work!"
}

func (u *userService) GetUserById(userId int) string {
	return fmt.Sprintf("User detail::%d\n", userId)
}

func (u *userService) UpdateUserById(userId int, payload string) string {
	return fmt.Sprintf("Update user works! userId=%d, payload=%s", userId, payload)
}
