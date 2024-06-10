package requests

import (
	"time"
)

type BookRequest struct {
	ArtistName string     `json:"artist_name" binding:"required"`
	StudioName string     `json:"studio_name" binding:"required"`
	Time       *time.Time `json:"time" binding:"required"`
}
