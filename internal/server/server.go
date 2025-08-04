package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

const (
	port         = "8080"
	shutdownTime = 5 * time.Second
)

func New(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}
}

func Start(ctx context.Context, log *slog.Logger, server *http.Server) error {
	go func() {
		<-ctx.Done()
		shutdownCtx, cancelCtx := context.WithTimeout(context.Background(), shutdownTime)
		defer cancelCtx()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.ErrorContext(shutdownCtx, "server shutdown failed", "error", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
