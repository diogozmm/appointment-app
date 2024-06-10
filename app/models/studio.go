package models

import (
	"time"

	"github.com/google/uuid"
)

type Studio struct {
	ID           uuid.UUID  `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Name         string     `gorm:"unique;not null" json:"name"`
	SocialNumber string     `json:"socialNumber"`
	Artists      []Artist   `gorm:"foreign_key:studio_Id"`
}
