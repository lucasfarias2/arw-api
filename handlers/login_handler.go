package handlers

import (
	"arw-api/models"
	"arw-api/services"
	"arw-api/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.FullUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Failed to decode request data: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	storedUser, err := services.GetFullUserByEmail(user.Email)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	isValid := utils.CheckPasswordHash(user.Password, storedUser.Password)

	if isValid {
		token, err := utils.GenerateToken(storedUser.ID, user.Email)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"token": token})
		if err != nil {
			log.Printf("Failed to encode response data: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
	}
}
