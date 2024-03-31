package websocketserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1021,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Recieved message:  %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func SetupServer(IPAddress string, port string) {
	http.HandleFunc("/websocket", websocketHandler)
	log.Fatal(http.ListenAndServe(IPAddress+port, nil))
}
