package services

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/repositories"
)

type AppointmentService interface {
	BookAppointment(requests.BookRequest) error
}

type appointmentService struct {
	appointmentRepository repositories.AppointmentRepository
}

// BookAppointment implements AppointmentService.
func (a *appointmentService) BookAppointment(request requests.BookRequest) error {
	appointment := models.NewAppointment(request)
	return a.appointmentRepository.Create(&appointment)
}

func InitAppointmentService(appointmentRepo repositories.AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepo}
}
