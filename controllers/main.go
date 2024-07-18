package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/migurd/waterwatch_back/services"
)

var db *sql.DB

type Controllers struct {
}

func New(dbConn *sql.DB) Controllers {
	db = dbConn
	return Controllers{}
}

func GetClaims(r *http.Request) (*services.Claims, error) {
	// Extract the JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, errors.New("bearer token missing")
	}

	// Validate the token and get claims
	claims, err := services.ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
