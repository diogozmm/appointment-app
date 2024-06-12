package repositories

import (
	"appointment-app/app/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type ArtistRepository interface {
	GetAllArtists() (*[]models.Artist, error)
	GetArtistByName(artistname string) (*models.Artist, error)
	GetAllArtistByStudio(studioId string) (*[]models.Artist, error)
	GetArtistByStudio(studioId string) (*models.Artist, error)

	SaveArtist(artist models.Artist) error
}

type artistRepository struct {
	db *gorm.DB
}

func (a *artistRepository) GetAllArtistByStudio(studioId string) (*[]models.Artist, error) {
	var studio models.Studio
	if err := a.db.Preload("Artists").First(&studio, studioId).Error; err != nil {
		log.Fatal("Error:", err)
		return nil, fmt.Errorf("Error fetching studio: %w", err)
	}

	return &studio.Artists, nil
}

func (a *artistRepository) GetAllArtists() (*[]models.Artist, error) {
	var artists []models.Artist
	if err := a.db.Find(&artists).Error; err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	return &artists, nil
}

func (a *artistRepository) GetArtistByName(artistname string) (*models.Artist, error) {
	var artist models.Artist

	if err := a.db.Where("name = ?", artistname).First(&artist).Error; err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *artistRepository) GetArtistByStudio(studioId string) (*models.Artist, error) {
	var artist models.Artist
	if err := a.db.Where("studio_id = ?", studioId).First(&artist).Error; err != nil {
		log.Fatal("Error:", err)
		return nil, fmt.Errorf("Error fetching artist: %w", err)
	}

	return &artist, nil
}

func (r *artistRepository) SaveArtist(artist models.Artist) error {
	return r.db.Create(artist).Error
}

func InitArtistRepository(db *gorm.DB) ArtistRepository {
	return &artistRepository{db}
}
