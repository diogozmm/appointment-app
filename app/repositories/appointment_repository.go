package repositories

import (
	"appointment-app/app/models"

	"github.com/jinzhu/gorm"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func InitAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db}
}

func (r *appointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}
