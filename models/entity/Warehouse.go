package entity

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name"`
	Location  string          `json:"location"`
	Capacity  int             `json:"capacity"`
	Stocks    []StockResponse `json:"stocks,omitempty" gorm:"foreignKey:WarehouseID"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt gorm.DeletedAt  `json:"-"`
}

type WarehouseResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
}

func (WarehouseResponse) TableName() string {
	return "warehouses"
}
