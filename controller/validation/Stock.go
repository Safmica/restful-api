package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
)

func StockValidation(stock *entity.Stock) error {
	var count int64
	var existingStock entity.Stock

	if stock.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if err := database.DB.Model(&entity.Product{}).Where("id = ?", stock.ProductID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("product_id does not exist")
	}

	count = 0

	if err := database.DB.Model(&entity.Warehouse{}).Where("id = ?", stock.WarehouseID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("warehouse_id does not exist")
	}

	if stock.Quantity < 0 {
		return errors.New("quantity must be greater than 0")
	}

	if stock.Quantity == 0 {
		return errors.New("quantity is required")
	}

	if err := database.DB.Where("product_id =? AND warehouse_id =?", stock.ProductID, stock.WarehouseID).First(&existingStock).Error; err == nil {
		if existingStock.ID != 0 {
			return errors.New("stock already exists")
		}
	}

	return nil
}
