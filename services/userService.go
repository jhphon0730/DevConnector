package services

import (
	"jhphon0730/DevConnector/models"
	"jhphon0730/DevConnector/database"
)

type UserService interface {
	GetUser(user_id string) (*models.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) GetUser(user_id string) (*models.User, error) {
	var user *models.User
	if result := database.DB.First(&user, user_id); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
