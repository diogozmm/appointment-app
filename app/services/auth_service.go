package services

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/repositories"
	"appointment-app/app/shared"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(requests.LoginRequest) models.TokenResponse
}

type authService struct {
	userRepository repositories.UserRepository
}

func InitAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (a *authService) Login(request requests.LoginRequest) models.TokenResponse {
	user, err := a.userRepository.FindByUsername(request.Username)

	if err != nil {
		return models.TokenResponse{Error: err.Error()}
	}

	if user.ID == uuid.Nil || !shared.CheckPasswordHash(request.Password, user.Password) {
		return models.TokenResponse{Error: "Password or username is invalid"}
	}
	token := shared.GenerateJWT(user.ID)

	return models.TokenResponse{Token: token}
}
