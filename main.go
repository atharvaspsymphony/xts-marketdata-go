package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	simpleapi "test/simpleapi" // Import the simpleapi package
	"time"

	"github.com/gorilla/websocket"
)

var (
	UserID = ""
	Token  = ""
)

const (
	appKey         = "caf0e727f0887e7a597911"
	secretKey      = "Pxrw554$zH"
	Source         = "WEBAPI"
	url            = "https://developers.symphonyfintech.in"
	XTSMessageCode = 1512
	BroadcastMode  = "Full"
)

// Define the login request payload
var loginPayload = simpleapi.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}

func socket() {

	//Login to get the token
	response, err := simpleapi.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	UserID = response.Result.UserID
	Token = response.Result.Token
	fmt.Println("Login Response-->", response.Result)

	//Subscribe to the instruments which you want to get datafeed on socket
	var subscribePayload = simpleapi.SubscribeRequest{
		Instruments: []simpleapi.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	SubscribeResponse, err := simpleapi.Subscribe(subscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("SubscribeResponse-->", SubscribeResponse.Result)

	//Socket
	connectionURL := fmt.Sprintf("wss://%s/apimarketdata/socket.io/?token=%s&userID=%s&publishFormat=JSON&broadcastMode=%s&transport=websocket&EIO=3",
		"developers.symphonyfintech.in", Token, UserID, BroadcastMode)

	fmt.Println("Connection URL -->", connectionURL)

	u, _, err := websocket.DefaultDialer.Dial(connectionURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		return
	}
	defer u.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := u.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)

		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := u.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt received. Closing connection...")
			err := u.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error writing close message:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
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

	// ClientConfig
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

	// SearchbyString
	const searchString string = "TCS"
	SearchbyStringResponse, err := simpleapi.SearchByString(searchString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Response-->", SearchbyStringResponse)

	// GetSeries
	var exchangeSegment = "1"
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

	//GetStrikePrices
	var GetStrikesParams = `{
		"exchangeSegment":"2",
		"series":"OPTIDX",
		"symbol":"NIFTY",
		"expiryDate": "29Aug2024",
		"optionType":"CE",
		"strikePrice":"25000"
	}`
	GetStrikesResponse, err := simpleapi.GetStrikePrices(GetStrikesParams)
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
	GetOptionTypeResponse, err := simpleapi.GetOptionType(GetOptionTypeParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetOptionTypeResponse-->", GetOptionTypeResponse)

	//GetIndexList
	exchangeSegment = "11"
	GetIndexListResponse, err := simpleapi.GetIndexList(exchangeSegment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetIndexListResponse-->", GetIndexListResponse)

	//Quote
	var quotePayload = simpleapi.QuoteRequest{
		Instruments: []simpleapi.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
		PublishFormat:  "JSON",
	}
	QuoteResponse, err := simpleapi.Quotes(quotePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("QuoteResponse-->", QuoteResponse)

	//Subscribe
	var subscribePayload = simpleapi.SubscribeRequest{
		Instruments: []simpleapi.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	SubscribeResponse, err := simpleapi.Subscribe(subscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("SubscribeResponse-->", SubscribeResponse.Result)

	//Unsubscribe
	var unsubscribePayload = simpleapi.SubscribeRequest{
		Instruments: []simpleapi.Instrument{
			{ExchangeSegment: 1, ExchangeInstrumentID: 26000},
		},
		XtsMessageCode: 1501,
	}
	UnsubscribeResponse, err := simpleapi.UnSubscribe(unsubscribePayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("UnsubscribeResponse-->", UnsubscribeResponse.Result)

	//Logout
	LogoutResponse, err := simpleapi.Logout()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("LogoutResponse-->", LogoutResponse)
}
