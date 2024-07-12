package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
)

func UserValidation(user *entity.User) error {
	var existingUser entity.User

	if user.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if user.Role == "" {
		return errors.New("role is required")
	}

	result := database.DB.Where("email =?", user.Email).First(&existingUser)
	if result.Error == nil {
		return errors.New("email has been used")
	}

	return nil
}

func UserUpdateValidation(requestUser *entity.UserUpdateResponse, user *entity.User) error {
	var existingUser entity.User

	if requestUser.ID != 0 {
		return errors.New("id is not allowed to be input")
	}

	if requestUser.Name != "" {
		user.Name = requestUser.Name
	}

	if requestUser.Email != "" {
		return errors.New("email is not allowed to be changed")
	}

	if requestUser.NewPassword != "" {
		result := database.DB.First(&existingUser, requestUser.Password)
		if result.Error == nil {
			return errors.New("wrong password")
		}
		user.Password = requestUser.NewPassword
	}

	return nil
}
