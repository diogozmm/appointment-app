package services_test

import (
	"errors"
	"testing"

	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	"appointment-app/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	hashedPassword := "hashedPassword"
	mockUtils.On("HashPassword", mock.Anything).Return(hashedPassword, nil)

	registerRequest := requests.RegisterRequest{
		Username: "testuser",
		Password: "password",
		UserType: 0,
	}

	mockUserRepository.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	err := userService.Register(registerRequest)

	assert.NoError(t, err)
	mockUserRepository.AssertExpectations(t)
}

func TestRegister_Failure(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	mockUtils.On("HashPassword", mock.Anything).Return("", errors.New("hash error"))

	registerRequest := requests.RegisterRequest{
		Username: "testuser",
		Password: "",
		UserType: 0,
	}

	err := userService.Register(registerRequest)

	assert.Error(t, err)
	assert.Equal(t, "hash error", err.Error())
	mockUserRepository.AssertNotCalled(t, "Create", mock.Anything)
}

func TestLogin_Success(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	hashedPassword := "hashedPassword"

	// Mocking HashPassword to return a hashed password
	mockUtils.On("HashPassword", mock.Anything).Return(hashedPassword, nil)

	registerRequest := requests.RegisterRequest{
		Username: "testuser",
		Password: "password",
		UserType: 0,
	}

	mockUserRepository.On("Create", mock.AnythingOfType("*models.User")).Return(nil)
	errorr := userService.Register(registerRequest)
	assert.NoError(t, errorr)

	user := &models.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	// Correctly mock CheckPasswordHash with the expected arguments
	mockUtils.On("CheckPasswordHash", "password", hashedPassword).Return(true)

	// Mock GenerateJWT to return a mock token
	mockUtils.On("GenerateJWT", mock.Anything).Return("mockToken")

	loginRequest := requests.LoginRequest{
		Username: "testuser",
		Password: "password",
	}

	mockUserRepository.On("FindByUsername", "testuser").Return(user, nil)

	token, err := userService.Login(loginRequest)

	assert.NoError(t, err)
	assert.Equal(t, "mockToken", token)
	mockUserRepository.AssertExpectations(t)
	mockUtils.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	hashedPassword := "hashedPassword"

	// Mocking HashPassword to return a hashed password
	mockUtils.On("HashPassword", mock.Anything).Return(hashedPassword, nil)

	registerRequest := requests.RegisterRequest{
		Username: "testuser",
		Password: "password",
		UserType: 0,
	}

	mockUserRepository.On("Create", mock.AnythingOfType("*models.User")).Return(nil)
	errorr := userService.Register(registerRequest)
	assert.NoError(t, errorr)

	user := &models.User{
		Username: "testuser",
		Password: "hashedPassword",
	}
	mockUtils.On("CheckPasswordHash", "wrongpassword", hashedPassword).Return(false)

	loginRequest := requests.LoginRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}

	mockUserRepository.On("FindByUsername", "testuser").Return(user, nil)

	token, err := userService.Login(loginRequest)

	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, "invalid credentials", err.Error())
	mockUserRepository.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockUserRepository := new(mocks.UserRepositoryMock)
	mockUtils := new(mocks.UtilsMock)
	userService := services.InitUserService(mockUserRepository, mockUtils)

	loginRequest := requests.LoginRequest{
		Username: "nonexistentuser",
		Password: "password",
	}

	mockUserRepository.On("FindByUsername", "nonexistentuser").Return((*models.User)(nil), errors.New("user not found"))

	token, err := userService.Login(loginRequest)

	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, "user not found", err.Error())
	mockUserRepository.AssertExpectations(t)
}
