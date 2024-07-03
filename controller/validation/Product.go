package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
)

func ProductValidation(product *entity.Product) error {
	var count int64
	var existingProduct entity.Product

	if product.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if product.Name == "" {
		return errors.New("name is required")
	}

	if product.Description == "" {
		return errors.New("description is required")
	}

	if product.Price < 0 {
		return errors.New("price must be greater than 0")
	}

	if product.Price == 0 {
		return errors.New("price is required")
	}

	if err := database.DB.Model(&entity.Category{}).Where("id = ?", product.CategoryID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("category_id does not exist")
	}

	if product.CategoryID == 0 {
		return errors.New("category_id is required")
	}

	if len(product.WarehouseID) == 0 {
		return errors.New("at least one warehouse_id is required")
	}

	if err := database.DB.Where("name =?", product.Name).Where("category_id =?", product.CategoryID).First(&existingProduct).Error; err == nil && existingProduct.ID != product.ID {
		return errors.New("product with the same name and category already exists")
	}

	return nil
}
