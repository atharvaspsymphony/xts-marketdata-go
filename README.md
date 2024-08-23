# xts-marketdata-Go

API Documentation for XTS-MarketData API can be found in the below link.

https://symphonyfintech.com/xts-market-data-front-end-api/

The XTS market data API provides developer, data-scientist, financial analyst and investor the market data with very low latency. It provides market data from various Indian electronic exchanges.

With the use of the socket.io library, the API has streaming capability and will push data notifications in a JSON format.

There is also an examples folder available which illustrates how to create a connection to XTS marketdata component in order to subscribe to real-time events. Please request for apikeys with Symphony Fintech developer support team to start integrating your application with XTS OEMS.

## Installation
Clone the Github repo in your working dir
```bash
https://github.com/atharvaspsymphony/golangTest/tree/master/api
```

## Usage
Access the /api dir directly.

```go
import (
	"fmt"
	marketdata "test/api"
)
```
Initailize all the required constants & variables. 

## Detailed explanation of API

### Login
call the login API to generate the token
POST /auth/login
```go
var loginPayload = marketdata.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}
response, err := marketdata.Login(url, loginPayload)
```
Once the token is generated you can call any api provided in the documentation.

### ClientConfig
GET /config/clientConfig
```go
clientResponse, err := marketdata.ClientConfig()
```

### Quote
POST /instruments/quotes
```go
var quotePayload = marketdata.QuoteRequest{
    Instruments: []marketdata.Instrument{
        {ExchangeSegment: 1, ExchangeInstrumentID: 26000},
    },
    XtsMessageCode: 1501,
    PublishFormat:  "JSON",
}
QuoteResponse, err := marketdata.Quotes(quotePayload)
```

### Subscription
POST /instruments/subscription
```go
var subscribePayload = marketdata.SubscribeRequest{
    Instruments: []marketdata.Instrument{
        {ExchangeSegment: 1, ExchangeInstrumentID: 26000},
    },
    XtsMessageCode: 1501,
}
SubscribeResponse, err := marketdata.Subscribe(subscribePayload)
```

### Unsubscription
POST /instruments/subscription
```go
var unsubscribePayload = marketdata.SubscribeRequest{
    Instruments: []marketdata.Instrument{
        {ExchangeSegment: 1, ExchangeInstrumentID: 26000},
    },
    XtsMessageCode: 1501,
}
UnsubscribeResponse, err := marketdata.UnSubscribe(unsubscribePayload)
```

### GetSeries
GET /instruments/instrument/series
```go
var exchangeSegment = "1"
GetSeriesResponse, err := marketdata.GetSeries(exchangeSegment)
```

### GetEquitySymbol
GET /instruments/instrument/symbol
```go
var equitySymbolParams = `{
    "exchangeSegment":"1",
    "series":"EQ",
    "symbol":"Reliance"}`

GetEquitySymboleResponse, err := marketdata.GetEquitySymbol(equitySymbolParams)
```

### GetExpiryDate
GET /instruments/instrument/expiryDate
```go
var GetExpiryParams = `{
    "exchangeSegment":"2",
    "series":"FUTIDX",
    "symbol":"NIFTY"
    }`

GetExpiryResponse, err := marketdata.GetExpiry(GetExpiryParams)
```

### GetFutureSymbol
GET /instruments/instrument/futureSymbol
```go
var GetFutureSymbolParams = `{
    "exchangeSegment":"2",
    "series":"FUTIDX",
    "symbol":"NIFTY",
    "expiryDate": "29Aug2024"
}`
GetFutureSymbolResponse, err := marketdata.GetFutureSymbol(GetFutureSymbolParams)
```

### GetOptionSymbol
GET /instruments/instrument/optionSymbol
```go
var GetOptionSymbolParams = `{
    "exchangeSegment":"2",
    "series":"OPTIDX",
    "symbol":"NIFTY",
    "expiryDate": "29Aug2024",
    "optionType":"CE",
    "strikePrice":"25000"
}`
GetOptionSymbolResponse, err := marketdata.GetOptionSymbol(GetOptionSymbolParams)
```

### GetStrikePrice
GET /instruments/instrument/strikePrice
```go
var GetStrikesParams = `{
    "exchangeSegment":"2",
    "series":"OPTIDX",
    "symbol":"NIFTY",
    "expiryDate": "29Aug2024",
    "optionType":"CE",
    "strikePrice":"25000"
}`
GetStrikesResponse, err := marketdata.GetStrikePrices(GetStrikesParams)
```

### GetOptionType
GET /instruments/instrument/optionType
```go
var GetOptionTypeParams = `{
    "exchangeSegment":"2",
    "series":"OPTIDX",
    "symbol":"NIFTY",
    "expiryDate": "29Aug2024",
    "optionType":"CE",
    "strikePrice":"25000"
}`
GetOptionTypeResponse, err := marketdata.GetOptionType(GetOptionTypeParams)
```

### IndexList
GET /instruments/indexlist
```go
exchangeSegment = "11"
GetIndexListResponse, err := marketdata.GetIndexList(exchangeSegment)
```

### Search Instruments by ID
POST /search/instrumentsbyid
```go
var searchPayload = marketdata.SearchRequest{
    Source: "WebAPI",
    Instruments: []marketdata.Instrument{
        {ExchangeSegment: 1, ExchangeInstrumentID: 26000},
    },
}
SearchbyIDResponse, err := marketdata.SearchByID(searchPayload)
```

### Search Instruments by String
GET /search/instruments
```go
const searchString string = "TCS"
SearchbyStringResponse, err := marketdata.SearchByString(searchString)
```

### OHLC
GET /instruments/ohlc
```go
var params = `{
    "exchangeSegment": "11",
    "exchangeInstrumentID": "532540",
    "startTime": "Aug 08 2024 090000",
    "endTime": "Aug 08 2024 153000",
    "compressionValue": "60"
}`

ohlcResponse, err := marketdata.GetOHLC(params)
```

## Instantiating the XtsMarketDataWS
This component provides functionality to access the socket related events. All real-time events can be registered via XtsMarketDataWS . After token is generated, you can access the socket component and instantiate the socket. Note that you will need to subscribe to instrument using Subscribe api [here](#subscription).

Note:- XTS MarketData WebSocket is based on "https://socket.io/" library. This library is available in most of the programming languages. In this package code generic webocket is used to make client connection for socket. This is just an example to make connection for websocket. For more reliable socket connectin you will need to use socket-io library.

```go
const(
    BroadcastMode  = "Full"
    url            = ""
)
response, err := marketdata.Login(url, loginPayload)
UserID = response.Result.UserID
Token = response.Result.Token
marketdata.Socket(Token, UserID, BroadcastMode)
```