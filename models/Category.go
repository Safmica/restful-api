package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}
