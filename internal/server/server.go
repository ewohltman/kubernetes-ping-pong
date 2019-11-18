package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const defaultPort = "8080"

func New(port string, mux http.Handler, errorLogger *log.Logger) *http.Server {
	if port == "" {
		port = defaultPort
	}

	return &http.Server{
		Addr:     ":" + port,
		Handler:  mux,
		ErrorLog: errorLogger,
	}
}

func Start(server *http.Server) chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			server.ErrorLog.Printf("HTTP server Shutdown: %v", err)
		}

		close(shutdown)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		server.ErrorLog.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	return shutdown
}
