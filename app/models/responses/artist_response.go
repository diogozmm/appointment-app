package responses

import (
	"appointment-app/app/models"
	"appointment-app/app/models/requests"
	"appointment-app/app/models/shared"
)

type ArtistResponse struct {
	ArtistName   string                  `json:"artist_name" binding:"required"`
	StudioName   string                  `json:"studio_name" binding:"required"`
	SocialNumber string                  `json:"social_number" binding:"required"`
	Address      requests.AddressRequest `json:"address"`
}

func NewArtistResponse(artist models.Artist) ArtistResponse {
	return ArtistResponse{
		ArtistName:   artist.Name,
		StudioName:   artist.Studio.Name,
		SocialNumber: artist.SocialNumber,
		Address:      shared.ToAddressRequest(artist.Address),
	}
}
