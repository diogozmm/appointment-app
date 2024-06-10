package models

import (
	"appointment-app/app/models/shared"
	"time"

	"github.com/google/uuid"
)

type Studio struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name         string     `gorm:"unique;not null" json:"name"`
	SocialNumber string     `json:"socialNumber"`
	AddressID    uuid.UUID  `json:"address_id"`
	Artists      []Artist   `gorm:"foreignKey:StudioID"`
	Address      shared.Address
}
