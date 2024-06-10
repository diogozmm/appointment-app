package shared

import (
	"appointment-app/app/models/requests"

	"github.com/google/uuid"
)

type Address struct {
	AddressId uuid.UUID `gorm:"primary_key" json:"address_id"`
	ArtistId  uuid.UUID `json:"artist_id"`
	Street    string    `json:"street_name"`
	Number    int32     `json:"street_number"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
}

func ToAddressRequest(address Address) requests.AddressRequest {
	return requests.AddressRequest{
		Street:  address.Street,
		Number:  address.Number,
		City:    address.City,
		State:   address.State,
		Country: address.Country,
	}
}
