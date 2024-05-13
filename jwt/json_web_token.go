package json_web_token

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	secret_key := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	json_web_token, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return json_web_token, err
}

func VerifyToken(s string) (*jwt.Token, error) {
	err := godotenv.Load("../.env")
	secret_key := os.Getenv("SECRET_KEY")

	if err != nil {
		panic(err.Error())
	}

	token, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unauthorized %v", t.Header["alg"])
		}

		return []byte(secret_key), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil

}

func DecodeToken(s string) (jwt.MapClaims, error) {
	token, err := VerifyToken(s)
	if err != nil {
		return nil, err
	}

	claims, isValid := token.Claims.(jwt.MapClaims)

	if !token.Valid && !isValid {
		return nil, err
	}

	return claims, nil

}
