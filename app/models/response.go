package models

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// TokenResponse represents a response containing a JWT token.
type TokenResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}
