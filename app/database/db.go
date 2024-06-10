package database

import (
	"appointment-app/app/models"
	"appointment-app/app/models/shared"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=appointmentApp password=223589 sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = DB.Debug()

	DB.AutoMigrate(
		&models.User{},
		&models.Artist{},
		&models.Appointment{},
		&models.Studio{},
		&shared.Address{},
		&models.Client{})

	// Manually add foreign key constraints
	createConstraints()
}

func createConstraints() {

	//Artist
	DB.Exec("ALTER TABLE artists ADD CONSTRAINT fk_artists_studio FOREIGN KEY (studio_id) REFERENCES studios(id) ON UPDATE CASCADE ON DELETE SET NULL;")
	DB.Exec("ALTER TABLE artists ADD CONSTRAINT fk_artists_address FOREIGN KEY (address_id) REFERENCES addresses(id) ON UPDATE CASCADE ON DELETE SET NULL;")

	//Client
	DB.Exec("ALTER TABLE clients ADD CONSTRAINT fk_clients_artist FOREIGN KEY (artist_id) REFERENCES artists(id) ON UPDATE CASCADE ON DELETE SET NULL;")

	//Studio
	DB.Exec("ALTER TABLE studios ADD CONSTRAINT fk_studios_address FOREIGN KEY (address_id) REFERENCES addresses(id) ON UPDATE CASCADE ON DELETE SET NULL;")

	//Users
	DB.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_artist FOREIGN KEY (artist_id) REFERENCES artists(id) ON UPDATE CASCADE ON DELETE SET NULL;")
	DB.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_client FOREIGN KEY (client_id) REFERENCES clients(id) ON UPDATE CASCADE ON DELETE SET NULL;")

}
