package simpleapi

var marketDataRoutes = map[string]interface{}{
	"auth.login":       "/apimarketdata/auth/login",
	"auth.logout":      "/apimarketdata/auth/logout",
	"clientconfig":     "/apimarketdata/config/clientConfig",
	"ohlc":             "/apimarketdata/instruments/ohlc",
	"quote":            "/apimarketdata/instruments/quotes",
	"subscribe":        "/apimarketdata/instruments/subscription",
	"search.id":        "/apimarketdata/search/instrumentsbyid",
	"search.string":    "/apimarketdata/search/instruments",
	"get.series":       "/apimarketdata/instruments/instrument/series",
	"get.symbol":       "/apimarketdata/instruments/instrument/symbol",
	"get.expiry":       "/apimarketdata/instruments/instrument/expiryDate",
	"get.futureSymbol": "/apimarketdata/instruments/instrument/futureSymbol",
	"get.optionSymbol": "/apimarketdata/instruments/instrument/optionSymbol",
	"get.strikes":      "/apimarketdata/instruments/instrument/strikePrice",
	"get.optionType":   "/apimarketdata/instruments/instrument/optionType",
	"get.indexlist":    "/apimarketdata/instruments/indexlist",
}
