package services

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/models/responses"
	"appointment-app/app/repositories"
)

type ArtistService interface {
	GetAllArtists() ([]responses.ArtistResponse, error)
	GetArtistByName(artistname string) (responses.ArtistResponse, error)
	GetAllArtistByStudio(studioId string) ([]responses.ArtistResponse, error)
	GetArtistByStudio(studioId string) (responses.ArtistResponse, error)

	RegisterArtist(request requests.ArtistRegisterRequest) (*responses.ArtistResponse, error)
}

type artistService struct {
	artistRepository repositories.ArtistRepository
}

func (a *artistService) GetAllArtistByStudio(studioId string) ([]responses.ArtistResponse, error) {
	artists, err := a.artistRepository.GetAllArtistByStudio(studioId)
	var result []responses.ArtistResponse
	if err != nil {
		return nil, err
	}

	for _, artist := range *artists {
		result = append(result, responses.NewArtistResponse(artist))
	}
	return result, nil
}

func (a *artistService) GetAllArtists() ([]responses.ArtistResponse, error) {
	artists := a.artistRepository.GetAllArtists()
	var result []responses.ArtistResponse

	for _, artist := range *artists {
		result = append(result, responses.NewArtistResponse(artist))
	}
	return result, nil
}

func (a *artistService) GetArtistByName(artistname string) (responses.ArtistResponse, error) {
	artist, err := a.artistRepository.GetArtistByName(artistname)

	if err != nil {
		return responses.ArtistResponse{}, err
	}

	return responses.NewArtistResponse(*artist), nil
}

func (a *artistService) GetArtistByStudio(studioId string) (responses.ArtistResponse, error) {
	artist, err := a.artistRepository.GetArtistByStudio(studioId)
	if err != nil {
		return responses.ArtistResponse{}, err
	}

	return responses.NewArtistResponse(*artist), nil
}

func (a *artistService) RegisterArtist(request requests.ArtistRegisterRequest) (*responses.ArtistResponse, error) {
	artist, err := a.artistRepository.GetArtistByName(request.Name)
	if err != nil || artist != nil {
		return &responses.ArtistResponse{}, err
	}

	model := models.ToArtistModel(request)

	isCreated := a.artistRepository.SaveArtist(model)
	if isCreated != nil {
		return nil, err
	}
	response := responses.NewArtistResponse(model)
	return &response, err
}

func InitArtistService(artistRepo repositories.ArtistRepository) ArtistService {
	return &artistService{artistRepo}
}
