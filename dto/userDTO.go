package dto 

import (
	"time"
)

type CreateUserDTO struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserExperienceDTO struct {
	Title       string      `json:"title"`
	Company     string      `json:"company"`
	Description string      `json:"description"`
  StartDate   time.Time   `json:"start_date"`
  EndDate     *time.Time  `json:"end_date"`
}

type UpdateUserExperienceDTO struct {
	Title       string      `json:"title"`
	Company     string      `json:"company"`
	Description string      `json:"description"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     *time.Time  `json:"end_date"`
}
