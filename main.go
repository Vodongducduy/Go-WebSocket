package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Hello World...")
	setupRouter()
	http.ListenAndServe(":8080", nil)
}

func reader(conn *websocket.Conn) {
	for {
		//read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("reader", err)
			return
		}

		//Print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//upgrader this connection to a WebSocket
	//Connection

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	errSendClient := ws.WriteMessage(1, []byte("Hi Client!"))
	if errSendClient != nil {
		log.Println("wsEndpoint: ", errSendClient)
	}

	reader(ws)
}

func setupRouter() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
