package config

import (
	"appointment-app/app/controllers"
)

type Controllers struct {
	ArtistController      *controllers.ArtistController
	UserController        *controllers.UserController
	AuthController        *controllers.AuthController
	AppointmentController *controllers.AppointmentController
}

func InitControllers(services Services) *Controllers {
	return &Controllers{
		ArtistController:      controllers.InitArtistController(services.ArtistService),
		UserController:        controllers.InitUserController(services.UserService),
		AuthController:        controllers.InitAuthController(services.AuthService),
		AppointmentController: controllers.InitAppointmentController(services.AppointmentService),
	}
}
