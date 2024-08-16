package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var stringKey = os.Getenv("MY_SECRET_KEY")
var jwtKey = []byte(stringKey)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(id int64, username string, role string) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour) // 1 day expiration
	claims := &Claims{
		ID:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        strconv.FormatInt(id, 10), // Store ID as a string
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Convert ID from string to int64
	id, err := strconv.ParseInt(claims.StandardClaims.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	claims.ID = id

	return claims, nil
}
