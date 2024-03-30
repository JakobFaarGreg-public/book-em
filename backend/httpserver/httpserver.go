package httpserver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"

	"backend/httpserver/model"

	_ "github.com/lib/pq"
)

var (
	IPAddress = "127.0.0.1"
	Port      = ":3333"
)

type key string

const keyServerAddr key = "serverAddr"

func getBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: /v1/books\n", ctx.Value(keyServerAddr))

	hasBook := r.URL.Query().Has("book")
	if !hasBook {
		_, err := io.WriteString(w, "/v1/books takes a book as a query argument\n")
		if err != nil {
			fmt.Printf("Error writing to response writer")
		}
		return
	}
	queryParam := r.URL.Query().Get("book")
	bookID, _ := strconv.Atoi(queryParam)
	fmt.Printf("/v1/books/%d\n", bookID)

	book, err := bookDAO.GetBookByID(bookID)
	if err != nil {
		fmt.Printf("Error calling GetBookByID")
	}
	fmt.Print(*book)
	_, err = io.WriteString(w, fmt.Sprintf("/v1/books/%d\n%s\n", bookID, book))
	if err != nil {
		fmt.Printf("Error writing to response writer")
	}
}

var bookDAO *model.BookDAO

func SetupServer() {
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

	connStr := "" // Get from env file
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error conecting to postgres database: %s", err)
	}
	defer db.Close()

	bookDAO = model.NewBookDAO(db)

	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		ExitFailedStatusCode := 1
		os.Exit(ExitFailedStatusCode)
	}
	cancelCtx()
}