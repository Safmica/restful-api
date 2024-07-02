package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	Name        string              `json:"name" form:"name"`
	Description string              `json:"description" form:"description"`
	Price       float64             `json:"price" form:"price"`
	CategoryID  uint                `json:"category_id" form:"category_id"`
	Category    CategoryResponse    `gorm:"foreignKey:CategoryID"`
	Stocks      []StockResponse     `json:"-" gorm:"foreignKey:ProductID"`
	WarehouseID []uint              `json:"warehouse_id,omitempty" form:"warehouse_id" gorm:"-"`
	Warehouses  []WarehouseResponse `json:"warehouses" gorm:"many2many:product_warehouses;foreignKey:ID;joinForeignKey:ProductID;References:ID;joinReferences:WarehouseID"`
	CreatedAt   time.Time           `json:"-"`
	UpdatedAt   time.Time           `json:"-"`
	DeletedAt   gorm.DeletedAt      `json:"-"`
}

type ProductResponseStock struct {
	ID          uint             `json:"-" gorm:"primaryKey"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	CategoryID  uint             `json:"category_id"`
	Category    CategoryResponse `json:"category"`
}

type ProductResponseCategory struct {
	ID          uint    `json:"-" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"-"`
}

func (ProductResponseStock) TableName() string {
	return "products"
}

func (ProductResponseCategory) TableName() string {
	return "products"
}
