package mocks

import (
	"appointment-app/app/models"

	"github.com/stretchr/testify/mock"
)

type ArtistRepositoryMock struct {
	mock.Mock
}

func (m *ArtistRepositoryMock) GetAllArtists() (*[]models.Artist, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Artist), args.Error(1)
}

func (m *ArtistRepositoryMock) GetArtistByName(name string) (*models.Artist, error) {
	args := m.Called(name)
	return args.Get(0).(*models.Artist), args.Error(1)
}

func (m *ArtistRepositoryMock) GetAllArtistByStudio(studioId string) (*[]models.Artist, error) {
	args := m.Called(studioId)
	return args.Get(0).(*[]models.Artist), args.Error(1)
}

func (m *ArtistRepositoryMock) GetArtistByStudio(studioId string) (*models.Artist, error) {
	args := m.Called(studioId)
	return args.Get(0).(*models.Artist), args.Error(1)
}

func (m *ArtistRepositoryMock) SaveArtist(artist models.Artist) error {
	args := m.Called(artist)
	return args.Error(0)
}
