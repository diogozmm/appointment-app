package config

import (
	"appointment-app/app/services"
)

type Services struct {
	ArtistService      services.ArtistService
	UserService        services.UserService
	AuthService        services.AuthService
	AppointmentService services.AppointmentService
}

func InitServices(repos Repositories, utils Utils) *Services {
	return &Services{
		ArtistService:      services.InitArtistService(repos.ArtistRepo),
		UserService:        services.InitUserService(repos.UserRepo, utils.Utils),
		AuthService:        services.InitAuthService(repos.UserRepo, utils.Utils),
		AppointmentService: services.InitAppointmentService(repos.AppointmentRepo),
	}
}
