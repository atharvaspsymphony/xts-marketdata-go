package socket

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func Socket(Token string, UserID string, BroadcastMode string) {

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
