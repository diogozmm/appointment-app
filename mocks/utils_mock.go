package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UtilsMock struct {
	mock.Mock
}

func (m *UtilsMock) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *UtilsMock) CheckPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

func (m *UtilsMock) GenerateJWT(userID uuid.UUID) string {
	args := m.Called(userID)
	return args.String(0)
}
