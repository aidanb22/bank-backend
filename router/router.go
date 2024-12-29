package router

import (
	"bank-app/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	//r.HandleFunc("/api/users", handlers.GetAllUsers).Methods("GET")
	//r.HandleFunc("/api/users/{id}", handlers.GetOneUser).Methods("GET")
	r.HandleFunc("/api/payments", handlers.MakePayment).Methods("POST")

	return r
}
