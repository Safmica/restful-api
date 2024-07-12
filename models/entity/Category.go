package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint                      `json:"id" gorm:"primaryKey"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Products    []ProductResponseCategory `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time                 `json:"-"`
	UpdatedAt   time.Time                 `json:"-"`
	DeletedAt   gorm.DeletedAt            `json:"-"`
}

type CategoryResponse struct {
	ID          uint           `json:"-"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}
