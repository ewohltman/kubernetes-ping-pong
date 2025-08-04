package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ewohltman/kubernetes-ping-pong/internal/handler"
	"github.com/ewohltman/kubernetes-ping-pong/internal/server"
)

const (
	pongServiceURL = "PONG_SERVICE_URL"
)

func run(log *slog.Logger) error {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	httpClient := &http.Client{Timeout: 5 * time.Second}

	psu := os.Getenv(pongServiceURL)
	if psu == "" {
		return fmt.Errorf("%s not defined", pongServiceURL)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Ping(log, httpClient, psu))

	return server.Start(ctx, log, server.New(mux))
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	if err := run(log); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
