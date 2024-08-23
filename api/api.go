package marketdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GenericHeader stores common headers.
var header = GenericHeader{
	contentType:   "application/json",
	Authorization: "",
}

// baseURL stores the base URL for the API.
var baseURL string = ""

// setHeaders sets common headers for a request.
func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", header.contentType)
	if header.Authorization != "" {
		req.Header.Set("Authorization", header.Authorization)
	}
}

// doRequest performs the HTTP request and returns the response body.
func doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// parseJSONResponse parses the JSON response into the provided interface.
func parseJSONResponse(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

// Login API.
func Login(url string, payload LoginRequest) (*LoginResponse, error) {
	baseURL = url

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+marketDataRoutes["auth.login"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response LoginResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	header.Authorization = response.Result.Token
	return &response, nil
}

// ClientConfig API.
func ClientConfig() (*GenericResponse, error) {
	req, err := http.NewRequest("GET", baseURL+marketDataRoutes["clientconfig"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetOHLC API.
func GetOHLC(params string) (*GenericResponse, error) {
	var queryParams map[string]string
	if err := json.Unmarshal([]byte(params), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["ohlc"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SearchByID API.
func SearchByID(payload SearchRequest) (*GenericResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+marketDataRoutes["search.id"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SearchByString API.
func SearchByString(searchString string) (*GenericResponse, error) {
	fullURL := fmt.Sprintf("%s?searchString=%s", baseURL+marketDataRoutes["search.string"].(string), searchString)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetSeries API
func GetSeries(exchangeSegment string) (*GenericResponse, error) {
	getSeriesURL := fmt.Sprintf("%s?exchangeSegment=%s", baseURL+marketDataRoutes["get.series"].(string), exchangeSegment)
	req, err := http.NewRequest("GET", getSeriesURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetEquitySymbol API
func GetEquitySymbol(exchangeSegmentParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(exchangeSegmentParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetEquitySymbolURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.symbol"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetEquitySymbolURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetExpiry API
func GetExpiry(GetExpiryParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(GetExpiryParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetExpiryURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.expiry"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetExpiryURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetFutureSymbol API
func GetFutureSymbol(GetFutureSymbolParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(GetFutureSymbolParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetFutureSymbolURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.futureSymbol"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetFutureSymbolURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetOptionSymbol API
func GetOptionSymbol(GetOptionSymbolParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(GetOptionSymbolParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetOptionSymbolURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.optionSymbol"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetOptionSymbolURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetStrikes API
func GetStrikePrices(GetStrikesParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(GetStrikesParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetStrikeURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.strikes"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetStrikeURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetOptionType API
func GetOptionType(GetOptionTypeParams string) (*GenericResponse, error) {

	var queryParams map[string]string
	if err := json.Unmarshal([]byte(GetOptionTypeParams), &queryParams); err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	for key, value := range queryParams {
		urlParams.Add(key, value)
	}

	GetOptionTypeURL := fmt.Sprintf("%s?%s", baseURL+marketDataRoutes["get.optionType"].(string), urlParams.Encode())
	req, err := http.NewRequest("GET", GetOptionTypeURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIndexList API
func GetIndexList(exchangeSegment string) (*GenericResponse, error) {
	GetIndexListURL := fmt.Sprintf("%s?exchangeSegment=%s", baseURL+marketDataRoutes["get.indexlist"].(string), exchangeSegment)
	req, err := http.NewRequest("GET", GetIndexListURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Quotes API
func Quotes(quotePayload QuoteRequest) (*GenericResponse, error) {
	jsonData, err := json.Marshal(quotePayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+marketDataRoutes["quote"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Subscribe API
func Subscribe(subscribePayload SubscribeRequest) (*SubscribeResponse, error) {
	jsonData, err := json.Marshal(subscribePayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+marketDataRoutes["subscribe"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result SubscribeResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	//Storing the subscribe response in memory data.
	for _, quote := range result.Result.QuotesList {
		exSegment := strconv.Itoa(quote.ExchangeSegment)
		exid := strconv.Itoa(quote.ExchangeInstrumentID)
		LoadInMemory(result.Code, exSegment, exid, result.Result.ListQuotes)
	}

	return &result, nil
}

// UnSubscribe API
func UnSubscribe(UnsubscribePayload SubscribeRequest) (*UnsubscribeResponse, error) {
	jsonData, err := json.Marshal(UnsubscribePayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+marketDataRoutes["subscribe"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result UnsubscribeResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Logout API
func Logout() (*GenericResponse, error) {
	req, err := http.NewRequest("DELETE", baseURL+marketDataRoutes["auth.logout"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)
	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
