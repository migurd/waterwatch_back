package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Response struct to store the response of the endpoint
type Response struct {
	IsValidToken bool   `json:"is_valid_token"`
	TypeOfUser   string `json:"type_of_user,omitempty"`
}

// Middleware to validate the token
func validateToken(r *http.Request) (*Claims, bool) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		log.Println("No Authorization header provided")
		return nil, false
	}

	// Check if the Authorization header is in the correct format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		log.Println("Authorization header is not in the correct format")
		return nil, false
	}

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate the token using the ValidateJWT function
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		return nil, false
	}

	// Return the claims and a true boolean indicating a valid token
	return claims, true
}

func GetTokenInfo(w http.ResponseWriter, r *http.Request) {
	claims, isValid := validateToken(r)

	var response Response

	if isValid {
		response.IsValidToken = true
		response.TypeOfUser = claims.Role
	} else {
		response.IsValidToken = false
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
