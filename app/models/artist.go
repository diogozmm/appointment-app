package models

import (
	"appointment-app/app/models/requests"
	"appointment-app/app/models/shared"
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
	Name         string         `gorm:"unique;not null" json:"name"`
	SocialNumber string         `json:"socialNumber"`
	StudioID     uuid.UUID      `gorm:"type:uuid" json:"studio_id"`
	AddressID    uuid.UUID      `gorm:"type:uuid" json:"address_id"`
	Studio       Studio         `json:"studio"`
	Address      shared.Address `json:"address"`
}

func NewArtist(name, socialnumber string, studio_id, user_id uuid.UUID) *Artist {
	return &Artist{
		Name:         name,
		SocialNumber: socialnumber,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}
}

func ToArtistModel(request requests.ArtistRegisterRequest) Artist {
	return Artist{
		Name:         request.Name,
		SocialNumber: request.SocialNumber,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
