package controllers

import (
	"encoding/json"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"
	"strconv"
)

func GetChannelsByUser(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var channels []models.Channel
	if err := database.DB.Find(&channels, "member_ids LIKE ?", "%"+strconv.Itoa(int(user.ID))+"%").Error; err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(channels); err != nil {
		http.Error(w, "Failed to encode channels", http.StatusInternalServerError)
	}
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var channel models.Channel
	err := json.NewDecoder(r.Body).Decode(&channel)
	if err != nil {
		http.Error(w, "Failed to decode channel info", http.StatusBadRequest)
		return
	}

	database.DB.Create(&channel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(channel); err != nil {
		http.Error(w, "Failed to encode channels", http.StatusInternalServerError)
	}
}
