package controllers

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	_ "appointment-app/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func InitUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.RegisterRequest true "User"
// @Success 200 {object} responses.RegisterResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /register [post]
func (f *UserController) Register(c *gin.Context) {
	var request requests.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "BadRequest"})
		return
	}

	notCreated := f.UserService.Register(request)
	if notCreated != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: notCreated.Error()})
		return
	}
	c.JSON(http.StatusOK, "User created successfully")
}

// FindAll godoc
// @Summary Find all users
// @Description Find all users in db
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} responses.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /users [get]
func (f *UserController) FindAll(c *gin.Context) {

	response, err := f.UserService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
