package marketdata

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"
	"github.com/gorilla/websocket"
	"encoding/json"
)

// Initialize logger
// var AppLogger *Logger // Global shared logger

// func init() {
// 	var err error
// 	AppLogger, err = NewLogger("app2.log")
// 	if err != nil {
// 		log.Fatal("Failed to initialize logger:", err)
// 	}
// }


func getPortAndWSType(rawurl string) (string, string, string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", "", "", err
	}
	port := u.Host
	wsType := "ws"
	if strings.HasPrefix(u.Scheme, "https") {
		wsType = "wss"
	}
	// Extract handshake path from the URL path
	handshakePath := u.Path
	if !strings.HasSuffix(handshakePath, "/socket.io") {
		handshakePath = strings.TrimRight(handshakePath, "/") + "/socket.io"
	}
	return port, wsType, handshakePath, nil
}

type EventHandler func(data string)

// global registry of handlers
var eventHandlers = make(map[string]EventHandler)

// Allow users to register handlers
func On(eventName string, handler EventHandler) {
	eventHandlers[eventName] = handler
}

func dispatchEvent(message string) {
	if !strings.HasPrefix(message, "42") {
		return
	}

	payload := strings.TrimPrefix(message, "42")

	// Match: ["eventName", {...json...}]
	re := regexp.MustCompile(`^\["([^"]+)",(.*)\]$`)
	matches := re.FindStringSubmatch(payload)
	if len(matches) < 3 {
		log.Println("Invalid event format:", payload)
		return
	}

	eventName := matches[1]
	eventData := strings.TrimSpace(matches[2])

	// Special handling for binary packets
	if eventName == "xts-binary-packet" {
		log.Println("xts-binary-packet event received – waiting for binary frame...")
		return
	}

	if handler, ok := eventHandlers[eventName]; ok {
		handler(eventData)
	} else {
		log.Println("No handler registered for event:", eventName)
	}
}

func Socket(url string, Token string, UserID string, BroadcastMode string) {
	port, wsType, handshakePath, err := getPortAndWSType(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//Socket
	connectionURL := fmt.Sprintf("%s://%s%s/?token=%s&userID=%s&publishFormat=JSON&broadcastMode=%s&transport=websocket&EIO=3",
		wsType, port, handshakePath, Token, UserID, BroadcastMode)

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
			messageType, message, err := u.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			if messageType == websocket.BinaryMessage {
				// Handle binary message for XTS binary packets
				onXTsBinaryPacket(message)
			} else {
				// Handle text messages
				dispatchEvent(string(message))
			}
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

func ApibinarymarketdataSocket(url string, Token string, UserID string, BroadcastMode string) {
    port, wsType, handshakePath, err := getPortAndWSType(url)
    if err != nil {
        fmt.Println("Error:", err)
    }
    //Socket
    connectionURL := fmt.Sprintf("%s://%s%s/?token=%s&userID=%s&publishFormat=JSON&broadcastMode=%s&transport=websocket&EIO=3",
        wsType, port, handshakePath, Token, UserID, BroadcastMode)
 
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

		expectBinary := false
		countMsg := 0

		for {
			_, message, err := u.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			msgStr := string(message)
			if strings.Contains(msgStr, `"xts-binary-packet"`) {
				expectBinary = true
				continue
			}

			if expectBinary {
				jsonData, err := onXTsBinaryPacket(message)
				if err != nil {
					log.Printf("Error processing binary packet: %v\n", err)
				} else {
					for _, obj := range jsonData {
						for eventName, eventValue := range obj {
							// Marshal single event as compact JSON
							dataBytes, err := json.Marshal(eventValue)
							if err != nil {
								log.Printf("Error marshaling event '%s': %v\n", eventName, err)
								continue
							}

							// AppLogger.Info("[%s] Decoded JSON: %s", eventName, string(dataBytes))

							// Dispatch event
							if handler, ok := eventHandlers[eventName]; ok {
								handler(string(dataBytes))
							} else {
								log.Printf("No handler registered for event: %s\n", eventName)
							}
						}
					}
				}

				expectBinary = false
				countMsg++
				continue
			}

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
 