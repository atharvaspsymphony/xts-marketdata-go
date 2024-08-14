package simpleapi

// LoginRequest represents the structure of the login request payload
type LoginRequest struct {
	SecretKey string `json:"secretKey"`
	AppKey    string `json:"appKey"`
	Source    string `json:"source"`
}

// Result represents the structure of the result object in the response
type LoginResult struct {
	Token                 string `json:"token"`
	UserID                string `json:"userID"`
	AppVersion            string `json:"appVersion"`
	ApplicationExpiryDate string `json:"application_expiry_date"`
}

// GenericResponse represents a generic structure of the response payload
type GenericResponse struct {
	Type        string      `json:"type"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Result      LoginResult `json:"result"`
}
