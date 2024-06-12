package controllers

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	_ "appointment-app/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func InitAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

// Login godoc
// @Summary Login a user
// @Description Login a user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.LoginRequest true "User"
// @Success 200 {object} models.TokenResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func (a *AuthController) Login(c *gin.Context) {
	var request requests.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Bad Request"})
		return
	}

	token := a.AuthService.Login(request)

	if token.Error != "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: token.Error})
	}

	c.JSON(http.StatusOK, token)
}
