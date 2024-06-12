package mocks

import (
	"appointment-app/app/models"

	"github.com/stretchr/testify/mock"
)

type AppointmentRepositoryMock struct {
	mock.Mock
}

func (m *AppointmentRepositoryMock) Create(appointment *models.Appointment) error {
	args := m.Called(appointment)
	return args.Error(0)
}
