package main

import (
	"appointment-app/app/database"
	"appointment-app/config"
	_ "appointment-app/docs" // Import your generated docs
	"appointment-app/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Appointment Booking API
// @version 1.0
// @description This is an API for booking appointments with artists.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	database.InitDB()
	utils := config.InitUtils()
	// Initiate repositories
	repos := config.InitRepositories(database.DB)

	// Initiate services
	services := config.InitServices(*repos, *utils)

	// Initiate controllers
	controllers := config.InitControllers(*services)

	r := gin.Default()
	routes.InitRoutes(r, *controllers)

	r.Run(":8080")
}
