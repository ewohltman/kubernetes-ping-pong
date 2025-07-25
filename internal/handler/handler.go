package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	ping           = "ping!"
	pong           = "pong!"
	pongServiceURL = "http://pong-istio.default.svc.cluster.local"
)

func Ping(log *slog.Logger, httpClient *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()
		}()

		log.Info(ping)

		req, err := http.NewRequest(http.MethodGet, pongServiceURL, http.NoBody)
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}()

		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("invalid response code: %d", resp.StatusCode)
			log.Error(err.Error())
			http.Error(w, err.Error(), resp.StatusCode)
			return
		}

		if _, err := w.Write(nil); err != nil {
			log.Error(err.Error())
		}
	}
}

func Pong(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()
		}()

		log.Info(pong)

		if _, err := w.Write(nil); err != nil {
			log.Error(err.Error())
		}
	}
}
