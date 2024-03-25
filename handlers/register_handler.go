package handlers

import (
	"arw-api/models"
	"arw-api/services"
	"arw-api/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.FullUser
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	createdUser, err := services.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}
