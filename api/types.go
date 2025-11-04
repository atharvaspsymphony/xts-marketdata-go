package marketdata

type GenericHeader struct {
	contentType   string
	Authorization string
}

// GenericResponse represents a generic structure of the response payload
type GenericResponse struct {
	Type        string      `json:"type"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
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

type QuoteRequest struct {
	Instruments    []Instrument `json:"instruments"`
	XtsMessageCode int          `json:"xtsMessageCode"`
	PublishFormat  string       `json:"publishFormat"`
}

type SubscribeRequest struct {
	Instruments    []Instrument `json:"instruments"`
	XtsMessageCode int          `json:"xtsMessageCode"`
}

type Instrument struct {
	ExchangeSegment      int `json:"exchangeSegment"`
	ExchangeInstrumentID int `json:"exchangeInstrumentID"`
}

type SubscribeResponse struct {
	GenericResponse
	Result struct {
		Mdp                        int      `json:"mdp"`
		QuotesList                 []Quote  `json:"quotesList"`
		ListQuotes                 []string `json:"listQuotes"`
		RemainingSubscriptionCount int      `json:"remaining_subscription_count"`
	} `json:"result"`
}

type UnsubscribeResponse struct {
	GenericResponse
	XtsMessageCode int
	Unsublist      []Instrument
}

type Quote struct {
	ExchangeSegment      int `json:"exchangeSegment"`
	ExchangeInstrumentID int `json:"exchangeInstrumentID"`
}

// Touchline represents touchline market data (1501)
type Touchline struct {
	MessageCode           uint16  `json:"messageCode"`
	ExchangeSegment       int16   `json:"exchangeSegment"`
	ExchangeInstrumentID  int32   `json:"exchangeInstrumentID"`
	BookType              int16   `json:"bookType"`
	MarketType            int16   `json:"marketType"`
	LastTradedPrice       float64 `json:"lastTradedPrice"`
	LastTradedQuantity    int32   `json:"lastTradedQuantity"`
	LastUpdateTime        int64   `json:"lastUpdateTime"`
	LastTradedTimestamp   int64   `json:"lastTradedTimestamp"`
	AverageTradedPrice    float64 `json:"averageTradedPrice"`
	VolumeTradedToday     int64   `json:"volumeTradedToday"`
	TotalBuyQuantity      int64   `json:"totalBuyQuantity"`
	TotalSellQuantity     int64   `json:"totalSellQuantity"`
	TotalTradedValue      float64 `json:"totalTradedValue"`
	OpenPrice             float64 `json:"openPrice"`
	HighPrice             float64 `json:"highPrice"`
	LowPrice              float64 `json:"lowPrice"`
	ClosePrice            float64 `json:"closePrice"`
}

// MarketDepthEvent represents market depth data (1502)
type MarketDepthEvent struct {
	MessageCode          uint16 `json:"messageCode"`
	ExchangeSegment      int16  `json:"exchangeSegment"`
	ExchangeInstrumentID int32  `json:"exchangeInstrumentID"`
	BookType             int16  `json:"bookType"`
	MarketType           int16  `json:"marketType"`
	BuySellIndicator     int8   `json:"buySellIndicator"`
	Quantity             int32  `json:"quantity"`
	Price                float64 `json:"price"`
	NumberOfOrders       int16  `json:"numberOfOrders"`
	BbBuySellIndicator   int8   `json:"bbBuySellIndicator"`
	BbQuantity           int32  `json:"bbQuantity"`
	BbPrice              float64 `json:"bbPrice"`
	BbNumberOfOrders     int16  `json:"bbNumberOfOrders"`
}

// OpenInterest represents open interest data (1510)
type OpenInterest struct {
	MessageCode           uint16  `json:"messageCode"`
	ExchangeSegment       int16   `json:"exchangeSegment"`
	ExchangeInstrumentID  int32   `json:"exchangeInstrumentID"`
	BookType              int16   `json:"bookType"`
	MarketType            int16   `json:"marketType"`
	OpenInterest          int64   `json:"openInterest"`
	Timestamp             int64   `json:"timestamp"`
	ChangeInOpenInterest  int64   `json:"changeInOpenInterest"`
}