package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "SECRET_TOKEN"

func GenerateJwt(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJwt(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyJwt(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if !isOk || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
