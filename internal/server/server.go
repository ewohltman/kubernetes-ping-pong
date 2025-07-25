package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

const port = "8080"

func New(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}
}

func Start(ctx context.Context, log *slog.Logger, server *http.Server) error {
	go func() {
		<-ctx.Done()
		shutdownCtx := context.Background()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.ErrorContext(shutdownCtx, err.Error())
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
