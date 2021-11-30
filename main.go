package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//Define a connection upgrader.
var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

//homepage router
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page")
}

//Listening permanently for the connection.
func reader(conn *websocket.Conn){
	//loop runs forever.
	for{
		//Read the message from the server
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		//Echo message to the client
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
//Websocket endpoint handler
func wsEndPoint(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "WebSocket Endpoint")

	//Allowing any connection through our app regardless of the origin.
	upgrader.CheckOrigin = func( r *http.Request) bool {return  true}

	//Error catching.
	ws, err := upgrader.Upgrade(w, r, nil )
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully connected")

	reader(ws)
}

//defining routes for our application.
func setUpRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}
func main(){
	fmt.Println("Go WebSockets...")

	setUpRoutes()

	//Init our server at a defined port
	log.Fatal(http.ListenAndServe(":8000", nil))
}
