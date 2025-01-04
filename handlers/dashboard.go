package handlers

import (
	"encoding/json"
	"net/http"
)

// DashboardData represents the structure of the dashboard response
type DashboardData struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	response := DashboardData{
		Message: "Welcome to your dashboard!",
		Data:    "Here is some protected data specific to you.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
