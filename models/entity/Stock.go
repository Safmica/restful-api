package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID          uint                 `json:"id" gorm:"primaryKey"`
	ProductID   uint                 `json:"product_id"`
	Product     ProductResponseStock `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	WarehouseID uint                 `json:"warehouse_id"`
	Warehouse   WarehouseResponse    `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Quantity    int                  `json:"quantity"`
	CreatedAt   time.Time            `json:"-"`
	UpdatedAt   time.Time            `json:"-"`
	DeletedAt   gorm.DeletedAt       `json:"-"`
}

type StockResponse struct {
	ID          uint                 `json:"-" gorm:"primaryKey"`
	WarehouseID uint                 `json:"-"`
	ProductID   uint                 `json:"-"`
	Product     ProductResponseStock `json:"product" gorm:"foreignKey:ProductID"`
	Quantity    int                  `json:"quantity"`
}

func (StockResponse) TableName() string {
	return "stocks"
}
