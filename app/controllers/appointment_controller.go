package controllers

import (
	"appointment-app/app/database"
	"appointment-app/app/models"
	"appointment-app/app/services"
	_ "appointment-app/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppointmentController struct {
	appointmentService services.AppointmentService
}

func InitAppointmentController(appointmentService services.AppointmentService) *AppointmentController {
	return &AppointmentController{appointmentService}
}

// BookAppointment godoc
// @Summary Book an appointment
// @Description Book an appointment with an artist
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Appointment"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /appointments [post]
func (ac *AppointmentController) BookAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	userId, _ := c.Get("userID")
	appointment.ID = userId.(uuid.UUID)
	database.DB.Create(&appointment)
	c.JSON(http.StatusOK, appointment)
}

// GetAppointments godoc
// @Summary Get all appointments
// @Description Get a list of all appointments for the logged-in user
// @Tags appointments
// @Produce json
// @Success 200 {array} models.Appointment
// @Security BearerAuth
// @Router /appointments [get]
func (ac *AppointmentController) GetAppointments(c *gin.Context) {
	userId, _ := c.Get("userId")
	var appointments []models.Appointment
	database.DB.Where("user_id = ?", userId).Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}
