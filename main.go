package main

import (
	"fmt"
	simpleapi "test/simpleapi" // Import the simpleapi package
)

const url string = "https://developers.symphonyfintech.in"
const secretKey string = "Pxrw554$zH"
const appKey string = "caf0e727f0887e7a597911"

var UserID = ""

// Define the login request payload
var loginPayload = simpleapi.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}

func main() {
	//Login
	response, err := simpleapi.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	UserID = response.Result.UserID
	fmt.Println("Login Response-->", response.Result)

	//ClientConfig
	clientResponse, err := simpleapi.ClientConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ClientConfig Response-->", clientResponse.Result)

	//OHLC
	var params = `{
		"exchangeSegment": "11",
		"exchangeInstrumentID": "532540",
		"startTime": "Aug 08 2024 090000",
		"endTime": "Aug 08 2024 153000",
		"compressionValue": "60"
	}`

	ohlcResponse, err := simpleapi.GetOHLC(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("OHLC Response-->", ohlcResponse.Result)

	//SearchbyID
	var searchPayload = simpleapi.SearchRequest{
		Source: "WebAPI",
		Instruments: []simpleapi.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
	}
	SearchbyIDResponse, err := simpleapi.SearchByID(searchPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Response-->", SearchbyIDResponse)

	//SearchbyString
	// const searchString string = "TCS"
	// SearchbyStringResponse, err := simpleapi.SearchByString(searchString)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Search Response-->", SearchbyStringResponse)

	//GetSeries
	const exchangeSegment string = "1"
	GetSeriesResponse, err := simpleapi.GetSeries(exchangeSegment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Response-->", GetSeriesResponse)

	//GetEquitySymbol
	var equitySymbolParams = `{
		"exchangeSegment":"1",
		"series":"EQ",
		"symbol":"Reliance"}`

	GetEquitySymboleResponse, err := simpleapi.GetEquitySymbol(equitySymbolParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetEquitySymboleResponse-->", GetEquitySymboleResponse)

	//GetExpiry
	var GetExpiryParams = `{
		"exchangeSegment":"2",
		"series":"FUTIDX",
		"symbol":"NIFTY"
		}`

	GetExpiryResponse, err := simpleapi.GetExpiry(GetExpiryParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetExpiryResponse-->", GetExpiryResponse)

	//GetFutureSymbol
	var GetFutureSymbolParams = `{
		"exchangeSegment":"2",
		"series":"FUTIDX",
		"symbol":"NIFTY",
		"expiryDate": "29Aug2024"
	}`
	GetFutureSymbolResponse, err := simpleapi.GetFutureSymbol(GetFutureSymbolParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetFutureSymbolResponse-->", GetFutureSymbolResponse)

	//GetOptionSymbol
	var GetOptionSymbolParams = `{
		"exchangeSegment":"2",
		"series":"OPTIDX",
		"symbol":"NIFTY",
		"expiryDate": "29Aug2024",
		"optionType":"CE",
		"strikePrice":"25000"
	}`
	GetOptionSymbolResponse, err := simpleapi.GetOptionSymbol(GetOptionSymbolParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetOptionSymbolResponse-->", GetOptionSymbolResponse)
}
