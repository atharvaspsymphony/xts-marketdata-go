package simpleapi

// LoginRequest represents the structure of the login request payload
type LoginRequest struct {
	SecretKey string `json:"secretKey"`
	AppKey    string `json:"appKey"`
	Source    string `json:"source"`
}

// GenericResponse represents a generic structure of the response payload
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
