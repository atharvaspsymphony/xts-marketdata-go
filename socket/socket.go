package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Define constants
const (
	apiKey            = "caf0e727f0887e7a597911"
	apiSecret         = "Pxrw554$zH"
	Source            = "WEBAPI"
	URL               = "https://developers.symphonyfintech.in"
	XTSMessageCode    = 1512
	BroadcastMode     = "Full"
	SocketIOPath      = "/apimarketdata/socket.io/"
	ContentType       = "application/json"
	ReconnectDelay    = 1 * time.Second
	ReconnectMax      = 50000 * time.Millisecond
	ReconnectMaxTries = 10
)

var instruments = []map[string]interface{}{
	{"exchangeSegment": 1, "exchangeInstrumentID": 26000}, // nsecm
}

type LoginResponse struct {
	Result struct {
		Token  string `json:"token"`
		UserID string `json:"userID"`
	} `json:"result"`
}

type SubscriptionResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		RemainingSubscriptionCount int           `json:"Remaining_Subscription_Count"`
		ListQuotes                 []interface{} `json:"listQuotes"`
	} `json:"result"`
	Type string `json:"type"`
}

func login() (*LoginResponse, error) {
	data := map[string]string{
		"secretKey": apiSecret,
		"appKey":    apiKey,
		"source":    Source,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", URL+"/apimarketdata/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}

	return &loginResp, nil
}

func subscribeInstruments(token string) (*SubscriptionResponse, error) {
	data := map[string]interface{}{
		"instruments":    instruments,
		"xtsMessageCode": XTSMessageCode,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", URL+"/apimarketdata/instruments/subscription", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var subResp SubscriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&subResp); err != nil {
		return nil, err
	}

	return &subResp, nil
}

type MDSocketClient struct {
	url           string
	token         string
	userID        string
	broadcastMode string
	connection    *websocket.Conn
}

func NewMDSocketClient(url, token, userID, broadcastMode string) *MDSocketClient {
	return &MDSocketClient{
		url:           url,
		token:         token,
		userID:        userID,
		broadcastMode: broadcastMode,
	}
}

func (c *MDSocketClient) connect() {
	connectionURL := "wss://" + "developers.symphonyfintech.in" + "/apimarketdata/socket.io/?token=" + c.token + "&userID=" + c.userID + "&publishFormat=JSON&broadcastMode=" + c.broadcastMode + "&transport=websocket&EIO=3"
	fmt.Println("connectionURL-->", connectionURL)
	var err error
	c.connection, _, err = websocket.DefaultDialer.Dial(connectionURL, nil)
	if err != nil {
		logrus.Fatalf("WebSocket connection failed: %v", err)
	}
	logrus.Info("Market Data Socket connected successfully!")

	c.listen()
}

func (c *MDSocketClient) listen() {
	defer c.connection.Close()

}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	// Login to API
	loginResp, err := login()
	if err != nil {
		logrus.Fatalf("Login failed: %v", err)
	}
	logrus.Infof("Login successful: %+v", loginResp)

	// Subscribe to instruments
	subResp, err := subscribeInstruments(loginResp.Result.Token)
	if err != nil {
		logrus.Fatalf("Subscription failed: %v", err)
	}
	logrus.Infof("Subscription successful: %+v", subResp)

	// Initialize and connect the WebSocket client
	client := NewMDSocketClient(URL, loginResp.Result.Token, loginResp.Result.UserID, BroadcastMode)
	client.connect()

	// Keeping the application alive
	select {}
}
