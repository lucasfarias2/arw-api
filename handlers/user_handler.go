package handlers

import (
	"arw-api/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failed to encode response data: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}
}
