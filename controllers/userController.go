package controllers

import (
	"jhphon0730/DevConnector/dto"
	"jhphon0730/DevConnector/models"
	"jhphon0730/DevConnector/services"
	"jhphon0730/DevConnector/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
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

	res := response.CreateResponse(http.StatusOK, user, nil, "User found successfully.")
	c.JSON(http.StatusOK, res)
	return
}

// TEST_CURL: curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "email@gmail.com", "password": "password"}'
func (u *userController) CreateUser(c *gin.Context) {
	var user dto.CreateUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		res := response.CreateResponse(http.StatusBadRequest, nil, err, "")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	newUser := models.User{}
	newUser.CreateUser(user.Name, user.Email, user.Password)

	createdUser, err := u.userService.CreateUser(&newUser)
	if err != nil {
		res := response.CreateResponse(http.StatusInternalServerError, nil, err, "")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.CreateResponse(http.StatusCreated, createdUser, nil, "User created successfully.")
	c.JSON(http.StatusOK, res)
	return
}
