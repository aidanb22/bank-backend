// middleware/auth.go
package middleware

import (
	"bank-app/utils" // Ensure this path matches your module path
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authenticating request for:", r.URL.Path)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Missing Authorization header")
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization header format")
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optionally, extract user information and set it in context
		// userID, ok := (*claims)["user_id"].(string)
		// if ok {
		//     ctx := context.WithValue(r.Context(), "userID", userID)
		//     next.ServeHTTP(w, r.WithContext(ctx))
		// } else {
		//     http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		//     return
		// }

		log.Println("Authentication successful for user_id:", (*claims)["user_id"])

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
