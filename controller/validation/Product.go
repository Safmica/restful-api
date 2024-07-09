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

func ProductUpdateValidation(requestProduct *entity.ProductResponse, Product *entity.Product) error {
	var count int64

	if requestProduct.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if requestProduct.Name != "" && requestProduct.CategoryID > 0 {
		count = 0
		database.DB.Model(&entity.Product{}).Where("name = ? AND category_id = ?", requestProduct.Name, requestProduct.CategoryID).Count(&count)
		if count > 0 {
			return errors.New("product with the same name and category_id already exists")
		}
		Product.Name = requestProduct.Name
		Product.CategoryID = requestProduct.CategoryID
	} else {
		if requestProduct.Name != "" {
			count = 0
			if err := database.DB.Model(&entity.Product{}).
				Where("name = ? AND category_id = ?", requestProduct.Name, Product.CategoryID).
				Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return errors.New("product with the same name and category_id  already exists")
			}
			Product.Name = requestProduct.Name
		}

		if requestProduct.CategoryID > 0 {
			count = 0
			if err := database.DB.Model(&entity.Product{}).
				Where("name = ? AND CategoryID  = ?", Product.Name, requestProduct.CategoryID).
				Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return errors.New("product with the same name and category_id  already exists")
			}
			Product.CategoryID = requestProduct.CategoryID
		}
	}

	if requestProduct.Description != "" {
		Product.Description = requestProduct.Description
	}

	if requestProduct.Price > 0 {
		Product.Price = requestProduct.Price
	}

	return nil
}
