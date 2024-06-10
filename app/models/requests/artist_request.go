package requests

import "github.com/google/uuid"

type ArtistRegisterRequest struct {
	UserId       uuid.UUID      `json:"user_id" binding:"required"`
	StudioId     uuid.UUID      `json:"studio_id"`
	Name         string         `json:"artist_name" binding:"required"`
	SocialNumber string         `json:"social_number" binding:"required"`
	Address      AddressRequest `json:"address" binding:"required"`
}
