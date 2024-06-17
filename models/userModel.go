package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    Name           string         `gorm:"size:100;not null" json:"name"`
    Email          string         `gorm:"size:100;not null;unique" json:"email"`
    Password       string         `gorm:"size:255;not null" json:"-"`
    ProfilePicture string         `gorm:"size:255" json:"profile_picture"`
    Skills         string         `gorm:"type:text" json:"skills"`
    Experience     []Experience   `gorm:"foreignKey:UserID" json:"experience"`
    Education      []Education    `gorm:"foreignKey:UserID" json:"education"`
    CreatedAt      time.Time      `json:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at"`
}

// 사용자의 경력 정보를 담고 있으며, 사용자와 일대다 관계
type Experience struct {
    ID          uint        `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID      uint        `gorm:"not null" json:"user_id"`
    Title       string      `gorm:"size:100;not null" json:"title"`
    Company     string      `gorm:"size:100;not null" json:"company"`
    StartDate   time.Time   `json:"start_date"`
    EndDate     *time.Time  `json:"end_date"`
    Description string      `gorm:"type:text" json:"description"`
}

// 사용자의 학력 정보를 담고 있으며, 사용자와 일대다 관계
type Education struct {
    ID           uint        `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID       uint        `gorm:"not null" json:"user_id"`
    School       string      `gorm:"size:100;not null" json:"school"`
    Degree       string      `gorm:"size:100;not null" json:"degree"`
    FieldOfStudy string      `gorm:"size:100;not null" json:"field_of_study"`
    StartDate    time.Time   `json:"start_date"`
    EndDate      *time.Time  `json:"end_date"`
    Description  string      `gorm:"type:text" json:"description"`
}

// 사용자 생성 전에 실행할 작업 정의
func (user *User) BeforeCreate(tx *gorm.DB) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    return nil
}

// 경력 생성 전에 실행할 작업 정의
func (experience *Experience) BeforeCreate(tx *gorm.DB) error {
    experience.StartDate = time.Now()
    if experience.EndDate == nil {
        experience.EndDate = new(time.Time)  // 초기화하여 nil 포인터 문제 해결
    }
    return nil
}

// 학력 생성 전에 실행할 작업 정의
func (education *Education) BeforeCreate(tx *gorm.DB) error {
    education.StartDate = time.Now()
    if education.EndDate == nil {
        education.EndDate = new(time.Time)  // 초기화하여 nil 포인터 문제 해결
    }
    return nil
}
