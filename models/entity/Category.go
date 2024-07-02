package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint                      `json:"id" gorm:"primaryKey"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Products    []ProductResponseCategory `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time                 `json:"-"`
	UpdatedAt   time.Time                 `json:"-"`
	DeletedAt   gorm.DeletedAt            `json:"-"`
}

type CategoryResponse struct {
	ID          uint   `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}
