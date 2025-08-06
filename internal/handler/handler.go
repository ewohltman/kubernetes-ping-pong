package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	ping = "ping!"
	pong = "pong!"
)

func Ping(log *slog.Logger, httpClient *http.Client, pongServiceURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()
		}()

		log.InfoContext(r.Context(), ping)

		req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, pongServiceURL, http.NoBody)
		if err != nil {
			log.ErrorContext(r.Context(), "failed to create pong request", "error", err)
			http.Error(w, fmt.Sprintf("failed to create pong request: %s", err), http.StatusInternalServerError)
			return
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			log.ErrorContext(r.Context(), "error performing pong request", "error", err)
			http.Error(w, fmt.Sprintf("error performing pong request: %s", err), http.StatusInternalServerError)
			return
		}

		defer func() {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}()

		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("invalid response code: %d", resp.StatusCode)
			log.ErrorContext(r.Context(), "error validating response", "error", err)
			http.Error(w, fmt.Sprintf("error validating response: %s", err), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte(ping)); err != nil {
			log.ErrorContext(r.Context(), "error writing response", "error", err)
		}
	}
}

func Pong(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()
		}()

		log.InfoContext(r.Context(), pong)

		if _, err := w.Write([]byte(pong)); err != nil {
			log.ErrorContext(r.Context(), "error writing response", "error", err)
		}
	}
}
