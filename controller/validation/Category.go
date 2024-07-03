package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
)

func CategoryValidation(category *entity.Category) error {
	var existingCategory entity.Category

	if category.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if category.Name == "" {
		return errors.New("name is required")
	}

	if category.Description == "" {
		return errors.New("description is required")
	}

	if err := database.DB.Where("name =?", category.Name).First(&existingCategory).Error; err == nil {
		if existingCategory.ID != 0 {
			return errors.New("category already exists")
		}
	}

	return nil
}
