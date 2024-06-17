// /server/models/user.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User represents the user model
type User struct {
    ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    Name           string    `gorm:"size:100;not null" json:"name"`
    Email          string    `gorm:"size:100;not null;unique" json:"email"`
    Password       string    `gorm:"size:255;not null" json:"-"`
    ProfilePicture string    `gorm:"size:255" json:"profile_picture"`
    Skills         string    `gorm:"type:text" json:"skills"`
    Experience     []Experience `gorm:"foreignKey:UserID" json:"experience"`
    Education      []Education  `gorm:"foreignKey:UserID" json:"education"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

// 사용자의 경력 정보를 담고 있으며, 사용자와 일대다 관계
type Experience struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    Title       string    `gorm:"size:100;not null" json:"title"`
    Company     string    `gorm:"size:100;not null" json:"company"`
    StartDate   time.Time `json:"start_date"`
    EndDate     *time.Time `json:"end_date"`
    Description string    `gorm:"type:text" json:"description"`
}

// 사용자의 학력 정보를 담고 있으며, 사용자와 일대다 관계
type Education struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    School      string    `gorm:"size:100;not null" json:"school"`
    Degree      string    `gorm:"size:100;not null" json:"degree"`
    FieldOfStudy string   `gorm:"size:100;not null" json:"field_of_study"`
    StartDate   time.Time `json:"start_date"`
    EndDate     *time.Time `json:"end_date"`
    Description string    `gorm:"type:text" json:"description"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
    user.ID = uuid.New()
    return nil
}

func (experience *Experience) BeforeCreate(tx *gorm.DB) error {
    experience.ID = uuid.New()
    return nil
}

func (education *Education) BeforeCreate(tx *gorm.DB) error {
    education.ID = uuid.New()
    return nil
}
