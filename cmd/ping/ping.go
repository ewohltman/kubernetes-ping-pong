package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ewohltman/kubernetes-ping-pong/internal/handler"
	"github.com/ewohltman/kubernetes-ping-pong/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	httpClient := &http.Client{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Ping(logger, httpClient))

	shutdown := server.Start(server.New(os.Getenv("PORT"), mux, logger))

	<-shutdown

	logger.Println()
}
