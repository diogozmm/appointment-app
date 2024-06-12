package mocks

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"

	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) Login(request requests.LoginRequest) models.TokenResponse {
	args := m.Called(request)
	return args.Get(0).(models.TokenResponse)
}
