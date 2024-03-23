package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	IPAddress = "127.0.0.1"
	Port      = ":3333"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /v1 request\n")
	_, err := io.WriteString(w, "This is /v1 root!\n")
	if err != nil {
		fmt.Printf("Error writing to response writer")
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /v1/book request\n")
	_, err := io.WriteString(w, "This is /v1/book!\n")
	if err != nil {
		fmt.Printf("Error writing to response writer")
	}
}

func main() {
	http.HandleFunc("/v1", getRoot)
	http.HandleFunc("/v1/book", getBook)
	err := http.ListenAndServe(IPAddress+Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		ExitFailedStatusCode := 1
		os.Exit(ExitFailedStatusCode)
	}
}