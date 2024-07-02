package main

import (
	_ "go-authentication/docs" // Import generated docs
	"go-authentication/src/database"
	"go-authentication/src/routes"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title MyApp API
// @version 1.0
// @description This is a sample server for MyApp.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and then your token.

func main() {
	// Initialize the database connection
	database.Connect()

	// Set up the router
	r := routes.SetupRouter()

	// Swagger documentation route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
