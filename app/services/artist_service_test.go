package services_test

import (
	"errors"
	"testing"

	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/services"
	"appointment-app/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllArtists_Success(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	artists := &[]models.Artist{
		{ID: uuid.New(), Name: "Artist1"},
		{ID: uuid.New(), Name: "Artist2"},
	}

	mockArtistRepo.On("GetAllArtists").Return(artists, nil)

	result, err := artistService.GetAllArtists()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockArtistRepo.AssertExpectations(t)
}

func TestGetArtistByName_Success(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	artist := &models.Artist{ID: uuid.New(), Name: "Artist1"}

	mockArtistRepo.On("GetArtistByName", "Artist1").Return(artist, nil)

	result, err := artistService.GetArtistByName("Artist1")

	assert.NoError(t, err)
	assert.Equal(t, "Artist1", result.ArtistName)
	mockArtistRepo.AssertExpectations(t)
}

func TestGetArtistByName_NotFound(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	mockArtistRepo.On("GetArtistByName", "NonExistentArtist").Return((*models.Artist)(nil), errors.New("artist not found"))

	result, err := artistService.GetArtistByName("NonExistentArtist")

	assert.Error(t, err)
	assert.Equal(t, "", result.ArtistName)
	mockArtistRepo.AssertExpectations(t)
}

func TestGetAllArtistByStudio_Success(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	artists := &[]models.Artist{
		{ID: uuid.New(), Name: "Artist1", StudioID: uuid.New()},
		{ID: uuid.New(), Name: "Artist2", StudioID: uuid.New()},
	}

	mockArtistRepo.On("GetAllArtistByStudio", "1").Return(artists, nil)

	result, err := artistService.GetAllArtistByStudio("1")

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockArtistRepo.AssertExpectations(t)
}

func TestGetArtistByStudio_Success(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	artist := &models.Artist{ID: uuid.New(), Name: "Artist1", StudioID: uuid.New()}

	mockArtistRepo.On("GetArtistByStudio", "1").Return(artist, nil)

	result, err := artistService.GetArtistByStudio("1")

	assert.NoError(t, err)
	assert.Equal(t, "Artist1", result.ArtistName)
	mockArtistRepo.AssertExpectations(t)
}

func TestRegisterArtist_Success(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)

	artistRegisterRequest := requests.ArtistRegisterRequest{
		Name:     "NewArtist",
		StudioId: uuid.New(),
	}

	mockArtistRepo.On("GetArtistByName", "NewArtist").Return((*models.Artist)(nil), nil)

	mockArtistRepo.On("SaveArtist", mock.AnythingOfType("models.Artist")).Return(nil)

	result, err := artistService.RegisterArtist(artistRegisterRequest)

	assert.NoError(t, err)
	assert.Equal(t, "NewArtist", result.ArtistName)
	mockArtistRepo.AssertExpectations(t)
}

func TestRegisterArtist_AlreadyExists(t *testing.T) {
	mockArtistRepo := new(mocks.ArtistRepositoryMock)
	artistService := services.InitArtistService(mockArtistRepo)
	sameId := uuid.New()
	existingArtist := &models.Artist{ID: uuid.New(), Name: "ExistingArtist", StudioID: sameId}

	artistRegisterRequest := requests.ArtistRegisterRequest{
		Name:     "ExistingArtist",
		StudioId: sameId,
	}

	mockArtistRepo.On("GetArtistByName", "ExistingArtist").Return(existingArtist, nil)

	result, err := artistService.RegisterArtist(artistRegisterRequest)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockArtistRepo.AssertExpectations(t)
}
