package requests

type AddressRequest struct {
	Street  string `json:"street_name"`
	Number  int32  `json:"street_number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
