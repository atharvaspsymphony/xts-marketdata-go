package main

import (
	"fmt"
	"test/simpleapi" // Import the simpleapi package
)

const url string = "http://103.181.209.198:11091/apimarketdata/auth/login"
const secretKey string = "Ntag704@B8"
const appKey string = "04b4dc313c5be9932b3917"

// Define the login request payload
var loginPayload = simpleapi.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}

func main() {
	// Call the Login function
	response, err := simpleapi.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", response.Result)
}
