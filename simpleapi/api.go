package simpleapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login performs a POST request to the login API.
func Login(url string, payload LoginRequest) (*GenericResponse, error) {
	// Marshal the request body to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into the GenericResponse struct
	var response GenericResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Other API functions like GET, PUT can be added similarly.
