package websocketserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend/httpserver/model"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1021,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var bookDAO *model.BookDAO

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

		bookID, _ := strconv.Atoi(string(message))
		fmt.Printf("/v1/books/%d\n", bookID)

		book, err := bookDAO.GetBookByID(bookID)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Recieved message:  %s", book)

		err = conn.WriteMessage(websocket.TextMessage, []byte(book.String()))
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func SetupServer(IPAddress string, port string) {
	http.HandleFunc("/websocket", websocketHandler)

	connStr := "" // Get from env file
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error connecting to postgres database: %s", err)
	}
	defer db.Close()
	bookDAO = model.NewBookDAO(db)
	log.Fatal(http.ListenAndServe(IPAddress+port, nil))
}
