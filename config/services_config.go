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

func InitServices(repos Repositories) *Services {
	return &Services{
		ArtistService:      services.InitArtistService(repos.ArtistRepo),
		UserService:        services.InitUserService(repos.UserRepo),
		AuthService:        services.InitAuthService(repos.UserRepo),
		AppointmentService: services.InitAppointmentService(repos.AppointmentRepo),
	}
}
