package main

import (
	"fmt"
	"sync"

	"backend/httpserver"
	"backend/websocketserver"
)

var (
	IPAddress     = "127.0.0.1"
	HTTPPort      = ":3333"
	WebsocketPort = ":3334"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		fmt.Printf("Starting HTTP Server\n")
		httpserver.SetupServer(IPAddress, HTTPPort)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		fmt.Printf("Starting Websocket Server\n")
		websocketserver.SetupServer(IPAddress, WebsocketPort)
	}()

	waitGroup.Wait()
}
