package responses

import "appointment-app/app/enum"

type RegisterResponse struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewRegisterResponse(username string, password string) RegisterResponse {
	return RegisterResponse{
		Username: username,
		Password: password,
	}
}

type UserResponse struct {
	Username string        `json:"username" binding:"required"`
	UserType enum.UserType `json:"user_type" binding:"required"`
}

func NewUserResponse(username string, userType enum.UserType) UserResponse {
	return UserResponse{
		Username: username,
		UserType: userType,
	}
}
