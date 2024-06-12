package mocks

import (
	"appointment-app/app/models"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock implementation of userRepository
type UserRepositoryMock struct {
	mock.Mock
}

// Create mocks the Create method of userRepository
func (m *UserRepositoryMock) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// FindByUsername mocks the FindByUsername method of userRepository
func (m *UserRepositoryMock) FindByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

// FindAll mocks the FindAll method of userRepository
func (m *UserRepositoryMock) FindAll() (*[]models.User, error) {
	args := m.Called()
	return args.Get(0).(*[]models.User), args.Error(1)
}
