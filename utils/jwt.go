package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "ftxdrainer"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("INVALID TOKEN METHOD")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Could not parse token")
		return 0, errors.New("could not parse token")
	}

	isTokenValid := token.Valid

	if !isTokenValid {
		return 0, errors.New("INVALID TOKEN")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, nil
	}

	userid := int64(claims["id"].(float64))

	return userid, nil
}
