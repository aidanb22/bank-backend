package main

import (
	"bank-app/database"
	"bank-app/middleware"
	"bank-app/router"
	"log"
	"net/http"
)

func main() {
	database.ConnectDB()
	r := router.InitRouter()
	corsHandler := middleware.NewCORSHandler(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
