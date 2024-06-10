package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID           uuid.UUID  `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Name         string     `gorm:"unique;not null" json:"name"`
	SocialNumber string     `gorm:"not null" json:"social_number"`
	Artists      []Artist   `gorm:"many2many:user_artists;"`
	UserId       uuid.UUID  `json:"user_id"`
}
