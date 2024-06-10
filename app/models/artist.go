package models

import (
	"appointment-app/app/models/requests"
	"appointment-app/app/models/shared"
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID           uuid.UUID  `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Name         string     `gorm:"unique;not null" json:"name"`
	SocialNumber string     `json:"socialNumber"`
	StudioId     uuid.UUID  `json:"studio_id"`
	UserId       uuid.UUID  `json:"user_id"`
	AddressId    uuid.UUID
	Studio       Studio         `gorm:"foreignkey:StudioId" json:"studio"`
	Address      shared.Address `gorm:"foreignkey:AddressId" json:"address"`
}

func NewArtist(name, socialnumber string, studio_id, user_id uuid.UUID) *Artist {
	return &Artist{
		Name:         name,
		SocialNumber: socialnumber,
		StudioId:     studio_id,
		UserId:       user_id,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}
}

func ToArtistModel(request requests.ArtistRegisterRequest) Artist {
	return Artist{
		Name:         request.Name,
		SocialNumber: request.SocialNumber,
		StudioId:     request.StudioId,
		UserId:       request.UserId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}
}
