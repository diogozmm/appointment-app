package routes

import (
	"appointment-app/app/middleware"
	"appointment-app/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // assuming swagger files package is imported
	ginSwagger "github.com/swaggo/gin-swagger" // assuming swagger files package is imported
)

func InitRoutes(
	router *gin.Engine,
	controllers config.Controllers,
) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	router.POST("/register", controllers.UserController.Register)
	router.POST("/login", controllers.AuthController.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuth())
	{
		// Users
		authorized.GET("/users", controllers.UserController.FindAll)

		// Artists
		authorized.GET("/artists", controllers.ArtistController.GetArtists)
		authorized.POST("/artists/register", controllers.ArtistController.RegisterArtist)

		// Appointments
		authorized.POST("/appointments", controllers.AppointmentController.BookAppointment)
		authorized.GET("/appointments", controllers.AppointmentController.GetAppointments)
	}
}
