package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var IPAddress = "127.0.0.1"
var Port = ":3333"

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /v1 request\n")
	io.WriteString(w, "This is /v1 root!\n")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /v1/book request\n")
	io.WriteString(w, "This is /v1/book!\n")
}

func main() {
	http.HandleFunc("/v1", getRoot)
	http.HandleFunc("/v1/book", getBook)
	err := http.ListenAndServe(Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		ExitFailedStatusCode := 1
		os.Exit(ExitFailedStatusCode)
	}
}
