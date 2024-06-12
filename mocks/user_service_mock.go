package mocks

import (
	"appointment-app/app/models/requests"
	"appointment-app/app/models/responses"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
	UserRepo UserRepositoryMock
}

func (s *UserServiceMock) Register(request requests.RegisterRequest) error {
	args := s.Called(request)
	return args.Get(0).(error)
}

func (s *UserServiceMock) Login(request requests.LoginRequest) (string, error) {
	args := s.Called(request)
	return args.Get(0).(string), args.Error(1)
}

func (s *UserServiceMock) FindAll() ([]responses.UserResponse, error) {
	args := s.Called()
	return args.Get(0).([]responses.UserResponse), args.Error(1)
}
