package controllers

import (
	"encoding/json"
	"fmt"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	// fmt.Println(email)

	var dbUser models.User
	if err := database.DB.Where("email = ?", email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbUser); err != nil {
		http.Error(w, "Failed to encode user info", http.StatusInternalServerError)
	}
}

func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user").(string)

	response := fmt.Sprintf("Hello, %s! This is your API.", userEmail)
	w.Write([]byte(response))
}
