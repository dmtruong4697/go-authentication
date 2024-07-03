package routes

import (
	"go-authentication/src/controllers"
	"go-authentication/src/middlewares"
	"net/http"

	_ "go-authentication/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Unprotected routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	// Example of a protected route
	api.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		email := middlewares.GetUserEmail(r)
		w.Write([]byte("Hello, " + email))
	}).Methods("GET")

	api.HandleFunc("/get-user-info", controllers.GetUserInfo).Methods("GET")

	return r
}
