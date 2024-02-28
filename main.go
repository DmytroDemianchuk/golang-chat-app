package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dmytrodemianchuk/golang-chat-app/pkg/httpserver"
	"github.com/dmytrodemianchuk/golang-chat-app/pkg/ws"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		fmt.Println("http server is starting on :8080")
		httpserver.StartHTTPServer()
	} else if *server == "websocket" {
		fmt.Println("websocket server is starting on :8081")
		ws.StartWebsocketServer()
	} else {
		fmt.Println("invalid server. Available server: http or websocket")
	}
}
