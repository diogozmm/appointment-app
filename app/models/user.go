package models

import (
	"appointment-app/app/enum"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `sql:"index" json:"deleted_at,omitempty"`
	Username  string        `gorm:"unique;not null" json:"username"`
	Password  string        `gorm:"not null" json:"password"`
	UserType  enum.UserType `json:"user_type"`
	ArtistID  uuid.UUID     `gorm:"type:uuid" json:"artist_id"`
	ClientID  uuid.UUID     `gorm:"type:uuid" json:"client_id"`
	Artist    Artist        `gorm:"foreign_key:user_id"`
	Client    Client        `gorm:"foreign_key:user_id"`
}

func NewUser(username string, password string, userType enum.UserType) *User {
	return &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		UserType:  userType,
	}
}
