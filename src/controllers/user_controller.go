package controllers

import (
	"encoding/json"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"
)

type UpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

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

// UpdateUserInfo godoc
// @Summary Update user information
// @Description Update user information based on the email
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UpdateUser true "User object"
// @Success 200 {object} models.User
// @Failure 401 {string} string "User not found"
// @Failure 500 {string} string "Failed to update user information"
// @Router /api/update-user-info [put]
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var dbUser models.User
	if err := database.DB.Where("email = ?", email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var updatedUser UpdateUser
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	dbUser.Name = updatedUser.Name
	dbUser.Password = updatedUser.Password

	if err := database.DB.Save(&dbUser).Error; err != nil {
		http.Error(w, "Failed to update user information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbUser); err != nil {
		http.Error(w, "Failed to encode user info", http.StatusInternalServerError)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	for i := range users {
		users[i].Password = ""
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}
