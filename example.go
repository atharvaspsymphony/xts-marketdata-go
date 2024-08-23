package main

import (
	"fmt"
	marketdata "test/api"
)

var (
	UserID = ""
	Token  = ""
)

const (
	appKey         = ""
	secretKey      = ""
	Source         = "WEBAPI"
	url            = ""
	XTSMessageCode = 1512
	BroadcastMode  = "Full"
)

// Define the login request payload
var loginPayload = marketdata.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}

func socketTest() {

	//Login to get the token
	response, err := marketdata.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	UserID = response.Result.UserID
	Token = response.Result.Token
	fmt.Println("Login Response-->", response.Result)

	//Subscribe to the instruments which you want to get datafeed on socket
	var subscribePayload = marketdata.SubscribeRequest{
		Instruments: []marketdata.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	SubscribeResponse, err := marketdata.Subscribe(subscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("SubscribeResponse-->", SubscribeResponse.Result)

	//Socket
	marketdata.Socket(Token, UserID, BroadcastMode)

}

func main() {
	//Login
	response, err := marketdata.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	UserID = response.Result.UserID
	fmt.Println("Login Response-->", response.Result)

	// ClientConfig
	clientResponse, err := marketdata.ClientConfig()
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

	ohlcResponse, err := marketdata.GetOHLC(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("OHLC Response-->", ohlcResponse.Result)

	//SearchbyID
	var searchPayload = marketdata.SearchRequest{
		Source: "WebAPI",
		Instruments: []marketdata.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
	}
	SearchbyIDResponse, err := marketdata.SearchByID(searchPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Response-->", SearchbyIDResponse)

	// SearchbyString
	const searchString string = "TCS"
	SearchbyStringResponse, err := marketdata.SearchByString(searchString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Response-->", SearchbyStringResponse)

	// GetSeries
	var exchangeSegment = "1"
	GetSeriesResponse, err := marketdata.GetSeries(exchangeSegment)
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

	GetEquitySymboleResponse, err := marketdata.GetEquitySymbol(equitySymbolParams)
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

	GetExpiryResponse, err := marketdata.GetExpiry(GetExpiryParams)
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
	GetFutureSymbolResponse, err := marketdata.GetFutureSymbol(GetFutureSymbolParams)
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
	GetOptionSymbolResponse, err := marketdata.GetOptionSymbol(GetOptionSymbolParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetOptionSymbolResponse-->", GetOptionSymbolResponse)

	//GetStrikePrices
	var GetStrikesParams = `{
		"exchangeSegment":"2",
		"series":"OPTIDX",
		"symbol":"NIFTY",
		"expiryDate": "29Aug2024",
		"optionType":"CE",
		"strikePrice":"25000"
	}`
	GetStrikesResponse, err := marketdata.GetStrikePrices(GetStrikesParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetStrikesResponse-->", GetStrikesResponse)

	//GetOptionType
	var GetOptionTypeParams = `{
		"exchangeSegment":"2",
		"series":"OPTIDX",
		"symbol":"NIFTY",
		"expiryDate": "29Aug2024",
		"optionType":"CE",
		"strikePrice":"25000"
	}`
	GetOptionTypeResponse, err := marketdata.GetOptionType(GetOptionTypeParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetOptionTypeResponse-->", GetOptionTypeResponse)

	//GetIndexList
	exchangeSegment = "11"
	GetIndexListResponse, err := marketdata.GetIndexList(exchangeSegment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetIndexListResponse-->", GetIndexListResponse)

	//Quote
	var quotePayload = marketdata.QuoteRequest{
		Instruments: []marketdata.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
		PublishFormat:  "JSON",
	}
	QuoteResponse, err := marketdata.Quotes(quotePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("QuoteResponse-->", QuoteResponse)

	//Subscribe
	var subscribePayload = marketdata.SubscribeRequest{
		Instruments: []marketdata.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	SubscribeResponse, err := marketdata.Subscribe(subscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("SubscribeResponse-->", SubscribeResponse.Result)

	//Unsubscribe
	var unsubscribePayload = marketdata.SubscribeRequest{
		Instruments: []marketdata.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	UnsubscribeResponse, err := marketdata.UnSubscribe(unsubscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("UnsubscribeResponse-->", UnsubscribeResponse.Result)

	//Logout
	LogoutResponse, err := marketdata.Logout()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("LogoutResponse-->", LogoutResponse)

	//socket test
	socketTest()
}
