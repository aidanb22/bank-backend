// middleware/cors.go
package middleware

import (
	"os"
	"strings"

	"github.com/rs/cors"
	"net/http"
)

func NewCORSHandler(next http.Handler) http.Handler {
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
	var allowedOrigins []string
	if allowedOriginsEnv != "" {
		allowedOrigins = strings.Split(allowedOriginsEnv, ",")
	} else {
		// Default to localhost:3000 if ALLOWED_ORIGINS is not set
		allowedOrigins = []string{"http://localhost:3000"}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	})

	return c.Handler(next)
}
