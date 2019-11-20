package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ewohltman/kubernetes-ping-pong/internal/handler"
	"github.com/ewohltman/kubernetes-ping-pong/internal/server"
)

const port = "30001"

func main() {
	logger := log.New(os.Stdout, "", 0)
	defer logger.Println()

	httpClient := &http.Client{Timeout: 5 * time.Second}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Ping(logger, httpClient))

	shutdown := server.Start(server.New(port, mux, logger))
	<-shutdown
}
