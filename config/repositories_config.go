package config

import (
	"appointment-app/app/repositories"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	ArtistRepo      repositories.ArtistRepository
	UserRepo        repositories.UserRepository
	AppointmentRepo repositories.AppointmentRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ArtistRepo:      repositories.InitArtistRepository(db),
		UserRepo:        repositories.InitUserRepository(db),
		AppointmentRepo: repositories.InitAppointmentRepository(db),
	}
}
