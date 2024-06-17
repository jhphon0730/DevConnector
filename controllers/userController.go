package controllers

import (
	"jhphon0730/DevConnector/services"
)

type UserController interface {

}

type userController struct {
	userService services.UserService
}

func NewUserController() UserController {
	return &userController{}
}
