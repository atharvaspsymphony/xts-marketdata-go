package main

import (
	"fmt"

	"../simpleapi"
)

func main() {
	// Define the login request payload
	loginPayload := simpleapi.LoginRequest{
		SecretKey: "Ntag704@B8",
		AppKey:    "04b4dc313c5be9932b3917",
		Source:    "WebAPI",
	}

	// Call the Login function
	url := "http://103.181.209.198:11091/apimarketdata/auth/login"
	response, err := simpleapi.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
