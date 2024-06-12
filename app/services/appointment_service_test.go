package services_test

import (
	"errors"
	"testing"
	"time"

	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	"appointment-app/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookAppointment_Success(t *testing.T) {
	mockAppointmentRepo := new(mocks.AppointmentRepositoryMock)
	appointmentService := services.InitAppointmentService(mockAppointmentRepo)
	today := time.Date(2024, 03, 01, 10, 4, 20, 0, time.UTC)
	bookRequest := requests.BookRequest{
		ArtistName: "Artist",
		StudioName: "Alcateia",
		Time:       &today,
	}

	// Mock the Create method to return nil, indicating success
	mockAppointmentRepo.On("Create", mock.AnythingOfType("*models.Appointment")).Return(nil)

	// Call the BookAppointment method
	err := appointmentService.BookAppointment(bookRequest)

	// Assert that no error was returned
	assert.NoError(t, err)

	// Assert that the Create method was called with the correct argument
	mockAppointmentRepo.AssertCalled(t, "Create", mock.AnythingOfType("*models.Appointment"))
}

func TestBookAppointment_Failure(t *testing.T) {
	mockAppointmentRepo := new(mocks.AppointmentRepositoryMock)
	appointmentService := services.InitAppointmentService(mockAppointmentRepo)

	today := time.Date(2024, 03, 01, 10, 4, 20, 0, time.UTC)
	bookRequest := requests.BookRequest{
		ArtistName: "Artist",
		StudioName: "Alcateia",
		Time:       &today,
	}

	// Mock the Create method to return an error
	mockAppointmentRepo.On("Create", mock.AnythingOfType("*models.Appointment")).Return(errors.New("create error"))

	// Call the BookAppointment method
	err := appointmentService.BookAppointment(bookRequest)

	// Assert that an error was returned
	assert.Error(t, err)
	assert.Equal(t, "create error", err.Error())

	// Assert that the Create method was called with the correct argument
	mockAppointmentRepo.AssertCalled(t, "Create", mock.AnythingOfType("*models.Appointment"))
}
