package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

var (
	IPAddress = "127.0.0.1"
	Port      = ":3333"
)

type key string

const keyServerAddr key = "serverAddr"

func getBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: /v1/book\n", ctx.Value(keyServerAddr))

	hasBook := r.URL.Query().Has("book")
	if hasBook {
		bookID := r.URL.Query().Get("book")
		fmt.Printf("/v1/books/%s\n", bookID)
		_, err := io.WriteString(w, fmt.Sprintf("/v1/books/%s\n", bookID))
		if err != nil {
			fmt.Printf("Error writing to response writer")
		}
	} else {
		fmt.Printf("/v1/books\n")
		_, err := io.WriteString(w, "/v1/books takes a book as a query argument\n")
		if err != nil {
			fmt.Printf("Error writing to response writer")
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/books", getBook)

	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    Port,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		ExitFailedStatusCode := 1
		os.Exit(ExitFailedStatusCode)
	}
	cancelCtx()
}
