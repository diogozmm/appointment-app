package services_test

import (
	"errors"
	"testing"

	"appointment-app/app/models"
	"appointment-app/app/models/responses"
	"appointment-app/app/services"
	"appointment-app/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindAll_Success(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	users := &[]models.User{
		{Username: "testuser1", UserType: 0},
		{Username: "testuser2", UserType: 0},
	}
	mockUserRepository.On("FindAll").Return(users, nil)

	expectedResponse := []responses.UserResponse{
		{Username: "testuser1", UserType: 0},
		{Username: "testuser2", UserType: 0},
	}

	result, err := userService.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	mockUserRepository.AssertExpectations(t)
}

func TestFindAll_Error(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	mockUserRepository.On("FindAll").Return((*[]models.User)(nil), errors.New("db error"))

	result, err := userService.FindAll()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "db error", err.Error())
	mockUserRepository.AssertExpectations(t)
}
