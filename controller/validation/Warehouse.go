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

func WarehouseUpdateValidation(requestWarehouse *entity.WarehouseResponse, warehouse *entity.Warehouse) error {
	var count int64

	if requestWarehouse.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if requestWarehouse.Name != "" && requestWarehouse.Location != "" {
		count = 0
		database.DB.Model(&entity.Warehouse{}).Where("name = ? AND location = ?", requestWarehouse.Name, requestWarehouse.Location).Count(&count)
		if count > 0 {
			return errors.New("warehouse with the same name and location already exists")
		}
		warehouse.Name = requestWarehouse.Name
		warehouse.Location = requestWarehouse.Location
	} else {
		if requestWarehouse.Name != "" {
			count = 0
			if err := database.DB.Model(&entity.Warehouse{}).
				Where("name = ? AND location = ?", requestWarehouse.Name, warehouse.Location).
				Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return errors.New("warehouse with the same name and location already exists")
			}
			warehouse.Name = requestWarehouse.Name
		}

		if requestWarehouse.Location != "" {
			count = 0
			if err := database.DB.Model(&entity.Warehouse{}).
				Where("name = ? AND location = ?", warehouse.Name, requestWarehouse.Location).
				Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return errors.New("warehouse with the same name and location already exists")
			}
			warehouse.Location = requestWarehouse.Location
		}
	}

	if requestWarehouse.Capacity > 0 {
		warehouse.Capacity = requestWarehouse.Capacity
	}

	return nil
}
