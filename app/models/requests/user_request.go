package requests

import "appointment-app/app/enum"

type RegisterRequest struct {
	Username string        `json:"username" binding:"required"`
	Password string        `json:"password" binding:"required"`
	UserType enum.UserType `json:"type" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
