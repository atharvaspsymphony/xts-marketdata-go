package marketdata

var marketDataRoutes = map[string]interface{}{
	"auth.login":       "/auth/login",
	"auth.logout":      "/auth/logout",
	"clientconfig":     "/config/clientConfig",
	"ohlc":             "/instruments/ohlc",
	"quote":            "/instruments/quotes",
	"subscribe":        "/instruments/subscription",
	"search.id":        "/search/instrumentsbyid",
	"search.string":    "/search/instruments",
	"get.series":       "/instruments/instrument/series",
	"get.symbol":       "/instruments/instrument/symbol",
	"get.expiry":       "/instruments/instrument/expiryDate",
	"get.futureSymbol": "/instruments/instrument/futureSymbol",
	"get.optionSymbol": "/instruments/instrument/optionSymbol",
	"get.strikes":      "/instruments/instrument/strikePrice",
	"get.optionType":   "/instruments/instrument/optionType",
	"get.indexlist":    "/instruments/indexlist",
}
