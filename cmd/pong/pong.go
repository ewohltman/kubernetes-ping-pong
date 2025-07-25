package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ewohltman/kubernetes-ping-pong/internal/handler"
	"github.com/ewohltman/kubernetes-ping-pong/internal/server"
)

func run(log *slog.Logger) error {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Pong(log))

	return server.Start(ctx, log, server.New(mux))
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	if err := run(log); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
