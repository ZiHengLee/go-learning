// client.go
package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func main() {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt)

	socketUrl := "wss://stream.openseabeta.com/socket/websocket?vsn=2.0.0&token=9e41bd366cce4d1d96eeb930da7a7f07"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl,nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)

	// Our main loop for the client
	// We send our relevant packets here
	//for {
	//	select {
	//	case <-time.After(time.Duration(1) * time.Millisecond * 1000):
	//		// Send an echo packet every second
	//		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from GolangDocs!"))
	//		if err != nil {
	//			log.Println("Error during writing to websocket:", err)
	//			return
	//		}
	//
	//	case <-interrupt:
	//		// We received a SIGINT (Ctrl + C). Terminate gracefully...
	//		log.Println("Received SIGINT interrupt signal. Closing all pending connections")
	//
	//		// Close our websocket connection
	//		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//		if err != nil {
	//			log.Println("Error during closing websocket:", err)
	//			return
	//		}
	//
	//		select {
	//		case <-done:
	//			log.Println("Receiver Channel Closed! Exiting....")
	//		case <-time.After(time.Duration(1) * time.Second):
	//			log.Println("Timeout in closing receiving channel. Exiting....")
	//		}
	//		return
	//	}
	//}
	time.Sleep(time.Second*60)
}
