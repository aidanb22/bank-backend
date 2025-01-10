package main

import (
	"bank-app/database"
	"bank-app/middleware"
	"bank-app/router"
	"bank-app/utils"           // Import the utils package
	"github.com/joho/godotenv" // Import godotenv
	"log"
	"net/http"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Continuing with environment variables.")
	}

	// Initialize the JWT key
	utils.InitializeJwtKey()

	// Connect to the database
	database.ConnectDB()

	// Initialize the router
	r := router.InitRouter()

	// Apply CORS middleware
	corsHandler := middleware.NewCORSHandler(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
