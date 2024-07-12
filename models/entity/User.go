package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"primaryKey"`
	Password    string         `json:"password"`
	NewPassword string         `json:"-" gorm:"-"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type UserResponse struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"primaryKey"`
	Role      string         `json:"role"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UserUpdateResponse struct {
	ID          uint           `json:"-" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"primaryKey"`
	Password    string         `json:"password"`
	NewPassword string         `json:"new_password" gorm:"-"`
	Role        string         `json:"role"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserUpdateResponse) TableName() string {
	return "users"
}
