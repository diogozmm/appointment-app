package controllers

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	_ "appointment-app/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	artistService services.ArtistService
}

func InitArtistController(artistService services.ArtistService) *ArtistController {
	return &ArtistController{artistService}
}

// GetArtists godoc
// @Summary Get all artists
// @Description Get a list of all artists
// @Tags artists
// @Produce json
// @Success 200 {array} models.Artist
// @Router /artists [get]
func (ac *ArtistController) GetArtists(c *gin.Context) {
	response, err := ac.artistService.GetAllArtists()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, response)
}

// RegisterArtist godoc
// @Summary Register an artist
// @Description Create an artist based on user
// @Tags artist-register
// @Produce json
// @Param artist body requests.ArtistRegisterRequest true "Artist"
// @Success 200 {object} responses.ArtistResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /artists/register [post]
func (ac *ArtistController) RegisterArtist(c *gin.Context) {
	var request requests.ArtistRegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Bad Request"})
		return
	}

	response, err := ac.artistService.RegisterArtist(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
