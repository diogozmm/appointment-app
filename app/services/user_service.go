package services

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/models/responses"
	"appointment-app/app/repositories"
	"appointment-app/app/shared"
	"errors"
)

type UserService interface {
	Register(requests.RegisterRequest) error
	Login(requests.LoginRequest) (string, error)
	FindAll() ([]responses.UserResponse, error)
}

type userService struct {
	userRepository repositories.UserRepository
	utils          shared.Utils
}

func InitUserService(userRepo repositories.UserRepository, utils shared.Utils) UserService {
	return &userService{userRepo, utils}
}

func (s *userService) Register(request requests.RegisterRequest) error {
	hashedPassword, err := s.utils.HashPassword(request.Password)
	if err != nil {
		return err
	}
	user := models.NewUser(request.Username, hashedPassword, request.UserType)
	return s.userRepository.Create(user)
}

func (s *userService) Login(request requests.LoginRequest) (string, error) {
	user, err := s.userRepository.FindByUsername(request.Username)
	if err != nil {
		return "", err
	}
	if !s.utils.CheckPasswordHash(request.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	return s.utils.GenerateJWT(user.ID), nil
}

func (s *userService) FindAll() ([]responses.UserResponse, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var result []responses.UserResponse

	for _, user := range *users {
		result = append(result, responses.NewUserResponse(user.Username, user.UserType))
	}
	return result, nil
}
