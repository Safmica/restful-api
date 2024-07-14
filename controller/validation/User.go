package validation

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"GDSC-PROJECT/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

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

func UserLoginValidation(user *entity.UserLogin) (string, error) {
	var existingUser entity.User

	if user.ID != 0 {
		return "", errors.New("id is not allowed to be input")
	}

	if user.Email == "" {
		return "", errors.New("email is required")
	}

	if user.Password == "" {
		return "", errors.New("password is required")
	}

	result := database.DB.Where("email =?", user.Email).First(&existingUser)
	if result.Error != nil {
		return "", errors.New("wrong credentials")
	}

	isValid := utils.VerifyPassword(existingUser.Password, user.Password)
	if !isValid {
		return "", errors.New("wrong credentials")
	}

	claims := jwt.MapClaims{
		"name":  existingUser.Name,
		"email": existingUser.Email,
		"role":  existingUser.Role,
		"exp":   jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
	}

	token, errGenerateToken := utils.GenerateJwt(&claims)
	if errGenerateToken != nil {
		return "", errGenerateToken
	}

	return token, nil
}
