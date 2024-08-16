package simpleapi

type GenericHeader struct {
	contentType   string
	Authorization string
}

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

type LoginResponse struct {
	GenericResponse
	Result LoginResult `json:"result"`
}

type SearchRequest struct {
	Source      string       `json:"source"`
	Instruments []Instrument `json:"instruments"`
}

type Instrument struct {
	ExchangeSegment      int `json:"exchangeSegment"`
	ExchangeInstrumentID int `json:"exchangeInstrumentID"`
}

// GenericResponse represents a generic structure of the response payload
type GenericResponse struct {
	Type        string      `json:"type"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}
