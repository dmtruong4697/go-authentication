package controllers

import (
	"fmt"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user").(string)

	response := fmt.Sprintf("Hello, %s! This is your API.", userEmail)
	w.Write([]byte(response))
}

func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user").(string)

	response := fmt.Sprintf("Hello, %s! This is your API.", userEmail)
	w.Write([]byte(response))
}
