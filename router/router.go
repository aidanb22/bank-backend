// router/router.go
package router

import (
	"bank-app/handlers"
	"bank-app/middleware"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Public Routes
	r.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/users/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/api/payments", handlers.MakePayment).Methods("POST")

	// Protected Routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware) // Apply AuthMiddleware

	protected.HandleFunc("/dashboard", handlers.GetDashboardData).Methods("GET")
	protected.HandleFunc("/account", handlers.GetAccountInfo).Methods("GET")
	protected.HandleFunc("/protected-route", handlers.ProtectedRoute).Methods("GET")

	return r
}
