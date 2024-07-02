package entity

type ProductWarehouse struct {
	ProductID   uint `json:"product_id" gorm:"primaryKey"`
	WarehouseID uint `json:"warehouse_id" gorm:"primaryKey"`
}

func (ProductWarehouse) TableName() string {
	return "product_warehouses"
}
