package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ewohltman/kubernetes-ping-pong/internal/handler"
	"github.com/ewohltman/kubernetes-ping-pong/internal/server"
)

const port = "8080"

func main() {
	logger := log.New(os.Stdout, "", 0)
	defer logger.Println()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Pong(logger))

	shutdown := server.Start(server.New(port, mux, logger))
	<-shutdown
}
