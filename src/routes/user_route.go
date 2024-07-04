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

	// auth routes
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

	// user route
	api.HandleFunc("/get-user-info", controllers.GetUserInfo).Methods("GET")
	api.HandleFunc("/update-user-info", controllers.UpdateUserInfo).Methods("PUT")
	api.HandleFunc("/all-user", controllers.GetAllUsers).Methods("POST")

	// chat route
	r.HandleFunc("/ws", controllers.HandleConnections).Methods("GET")
	r.HandleFunc("/messages", controllers.GetMessages).Methods("POST")

	// channel route
	api.HandleFunc("/channels", controllers.GetChannelsByUser).Methods("POST")
	api.HandleFunc("/create-channels", controllers.CreateChannel).Methods("POST")

	return r
}
