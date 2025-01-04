// utils/jwt.go
package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtKey is the secret key used to sign JWTs.
// It is exported to be accessible in other packages like middleware.
var JwtKey []byte

// Claims defines the structure of JWT claims.
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// InitializeJwtKey loads the JWT secret key from environment variables.
// This function should be called during application startup.
func InitializeJwtKey() {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}
	JwtKey = []byte(secret)
}

// GenerateJWT generates a JWT token for a given user ID.
func GenerateJWT(userID string) (string, error) {
	// Set token expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the user ID and expiry time
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "bank-app", // Replace with your app's name or domain
			Subject:   "user authentication",
		},
	}

	// Declare the token with the algorithm used for signing and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT parses and validates a JWT token string.
// It returns the claims if the token is valid.
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token with the claims
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and claims
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
