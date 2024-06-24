package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Capacity int     `json:"capacity"`
	Stocks   []Stock `gorm:"foreignKey:WarehouseID"`
}
