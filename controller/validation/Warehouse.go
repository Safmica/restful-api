package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
)

func WarehouseValidation(warehouse *entity.Warehouse) error {
	var existingWarehouse entity.Warehouse

	if warehouse.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if warehouse.Name == "" {
		return errors.New("name is required")
	}

	if warehouse.Location == "" {
		return errors.New("location is required")
	}

	if warehouse.Capacity < 0 {
		return errors.New("capacity must be greater than 0")
	}

	if warehouse.Capacity == 0 {
		return errors.New("capacity is required")
	}

	result := database.DB.Where("name = ? AND location = ?", warehouse.Name, warehouse.Location).First(&existingWarehouse)
	if result.Error == nil {
		return errors.New("warehouse with the same name and location already exists")
	}

	return nil
}
