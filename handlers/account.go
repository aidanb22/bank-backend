package handlers

import (
	"bank-app/middleware"
	"encoding/json"
	"log"
	"net/http"
	// Import your database package and models as needed
)

// AccountInfo represents the account details
type AccountInfo struct {
	ID            string  `json:"id"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// GetAccountInfo retrieves account information for the authenticated user
func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	log.Printf("Fetching account info for user ID: %s", userID)

	// Fetch account info from the database based on userID
	// Replace this mock data with actual database retrieval logic
	account := AccountInfo{
		ID:            userID,
		AccountNumber: "ACC123456789",
		Balance:       2500.50,
		CreatedAt:     "2023-01-01T00:00:00Z",
		UpdatedAt:     "2023-10-01T00:00:00Z",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		log.Println("Error encoding account info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
