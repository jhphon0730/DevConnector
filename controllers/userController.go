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

	CreateUserExperience(c *gin.Context)
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

// TEST_CURL: curl -X POST http://localhost:8080/api/users/ -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "email@gmail.com", "password": "password"}'
func (u *userController) CreateUser(c *gin.Context) {
	var unCreatedUser dto.CreateUserDTO
	if err := c.ShouldBindJSON(&unCreatedUser ); err != nil {
		res := response.CreateResponse(http.StatusBadRequest, nil, err, "")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	newUser := models.User{
		Email: unCreatedUser .Email,
		Name: unCreatedUser.Name,
		Password: unCreatedUser.Password,
	}

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

// TEST_CURL: curl -X POST http://localhost:8080/api/users/experience/1 -H "Content-Type: application/json" -d '[{"title": "Software Engineer", "company": "Google", "startDate": "2021-01-01", "endDate": "2021-12-31", "description": "Software Engineer at Google"}]'
func (u *userController) CreateUserExperience(c *gin.Context) {
	var unCreatedUserExperience []dto.CreateUserExperienceDTO
	if err := c.ShouldBindJSON(&unCreatedUserExperience); err != nil {
		res := response.CreateResponse(http.StatusBadRequest, nil, err, "")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := u.userService.GetUser(c.Param("user_id"))
	if err != nil {
		res := response.CreateResponse(http.StatusInternalServerError, nil, err, "")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	var experiences []models.Experience
	for _, experience := range unCreatedUserExperience {
		experience := models.Experience{
			UserID: user.ID,
			Title: experience.Title,
			Company: experience.Company,
			StartDate: experience.StartDate,
			EndDate: experience.EndDate,
			Description: experience.Description,
		}
		experiences = append(experiences, experience)
	}
	user.Experience = experiences

	err = u.userService.CreateUserExperience(user)
	if err != nil {
		res := response.CreateResponse(http.StatusInternalServerError, nil, err, "")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.CreateResponse(http.StatusCreated, user, nil, "Experience created successfully.")
	c.JSON(http.StatusCreated, res)
}
