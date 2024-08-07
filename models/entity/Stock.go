package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID          uint                   `json:"id" gorm:"primaryKey"`
	ProductID   uint                   `json:"product_id"`
	Product     ProductResponse        `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	WarehouseID uint                   `json:"warehouse_id"`
	Warehouse   WarehouseResponseStock `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Quantity    int                    `json:"quantity"`
	CreatedAt   time.Time              `json:"-"`
	UpdatedAt   time.Time              `json:"-"`
	DeletedAt   gorm.DeletedAt         `json:"-"`
}

type StockResponse struct {
	ID          uint            `json:"-" gorm:"primaryKey"`
	WarehouseID uint            `json:"-"`
	ProductID   uint            `json:"-"`
	Product     ProductResponse `json:"product" gorm:"foreignKey:ProductID"`
	Quantity    int             `json:"quantity"`
	DeletedAt   gorm.DeletedAt  `json:"-"`
}

func (StockResponse) TableName() string {
	return "stocks"
}
