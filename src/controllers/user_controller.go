package controllers

import (
	"encoding/json"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"
)

// GetUserInfo godoc
// @Summary Get user information
// @Description Get user information based on the email
// @Tags Users
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {string} string "User not found"
// @Failure 500 {string} string "Failed to encode user info"
// @Router /api/get-user-info [get]
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
