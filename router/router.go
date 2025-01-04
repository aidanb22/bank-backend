// Rename directory from middlewares to middleware

// router/router.go
package router

import (
	"bank-app/handlers"
	"bank-app/middleware" // Updated import path
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// User and Auth Routes
	r.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/users/login", handlers.LoginUser).Methods("POST")

	// Banking and Payment Routes
	r.HandleFunc("/api/payments", handlers.MakePayment).Methods("POST")

	// Protected Routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware) // Apply the AuthMiddleware

	// Add Dashboard Route
	protected.HandleFunc("/dashboard", handlers.GetDashboardData).Methods("GET")

	// Existing Protected Route (optional)
	protected.HandleFunc("/protected-route", handlers.ProtectedRoute).Methods("GET")

	return r
}
