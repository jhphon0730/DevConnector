package controllers

import (
	"jhphon0730/DevConnector/services"
	"jhphon0730/DevConnector/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	GetUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) GetUser(c *gin.Context) {
	user_id := c.Param("user_id")

	user, err := u.userService.GetUser(user_id)
	if err != nil {
		res := response.CreateResponse(http.StatusInternalServerError, nil, err, "")
		c.JSON(http.StatusInternalServerError, res)
		return
	
	}

	c.JSON(http.StatusOK, user)
	return
}
