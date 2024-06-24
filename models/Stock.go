package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductID   uint      `json:"product_id"`
	WarehouseID uint      `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	Product     Product   `gorm:"foreignKey:ProductID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
}
