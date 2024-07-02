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
	Stocks    []StockResponse `json:"stocks" gorm:"foreignKey:WarehouseID"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt gorm.DeletedAt  `json:"-"`
}

type WarehouseResponse struct {
	ID       uint   `json:"-" gorm:"primaryKey"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

/*
	type WarehouseResponseProduct struct {
		ID       uint    `json:"id" gorm:"primaryKey"`
		Name     string  `json:"name"`
		Location string  `json:"location"`
		Capacity int     `json:"capacity"`
		Stocks   []Stock `json:"stocks"`
	}
*/
func (WarehouseResponse) TableName() string {
	return "warehouses"
}

/*func (WarehouseResponseProduct) TableName() string {
	return "warehouses"
}*/
