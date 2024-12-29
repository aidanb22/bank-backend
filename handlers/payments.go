package handlers

import (
	"encoding/json"
	"net/http"
)

type Payment struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Mock response
	payment.ID = "67890"
	payment.Timestamp = "2024-12-28T12:00:00Z"
	json.NewEncoder(w).Encode(payment)
}
