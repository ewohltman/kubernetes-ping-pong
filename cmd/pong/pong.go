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

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Pong(logger))

	shutdown := server.Start(server.New(os.Getenv("PORT"), mux, logger))

	<-shutdown

	logger.Println()
}
