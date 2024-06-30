package services

import (
	"jhphon0730/DevConnector/models"
	"jhphon0730/DevConnector/database"
)

type UserService interface {
	GetUser(user_id string) (*models.User, error)
	GetUserWithExperience(user_id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	CreateUserExperience(user *models.User) (error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) GetUser(user_id string) (*models.User, error) {
	var user models.User
	if result := database.DB.First(&user, user_id); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *userService) GetUserWithExperience(user_id string) (*models.User, error) {
	var user models.User
	if result := database.DB.Preload("Experience").First(&user, user_id); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	if result := database.DB.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userService) CreateUserExperience(user *models.User) (error) {
	if result := database.DB.Save(user); result.Error != nil {
		return result.Error
	}

	return nil
}
