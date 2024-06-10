package models

import (
	"appointment-app/app/models/requests"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	ArtistID  uint       `json:"artist_id"`
	Time      string     `gorm:"not null" json:"time"`
}

func NewAppointment(requests.BookRequest) Appointment {
	return Appointment{}
}
