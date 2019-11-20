package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	ping           = "Ping!"
	pong           = "Pong!"
	pongServiceURL = "http://pong.default.svc.cluster.local:30002"
)

func Ping(logger *log.Logger, httpClient *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Print(ping)

		defer func() {
			err := r.Body.Close()
			if err != nil {
				logger.Printf("error closing ping request body: %v", err)
			}
		}()

		_, err := io.Copy(ioutil.Discard, r.Body)
		if err != nil {
			err = fmt.Errorf("error discarding ping request body: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf := bytes.NewBufferString(ping)

		pongReq, err := http.NewRequest(http.MethodPost, pongServiceURL, buf)
		if err != nil {
			err = fmt.Errorf("error creating pong request: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pongResp, err := httpClient.Do(pongReq)
		if err != nil {
			err = fmt.Errorf("error performing pong request: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() {
			err := pongResp.Body.Close()
			if err != nil {
				logger.Printf("Error closing pong response body: %v", err)
			}
		}()

		pongBody, err := ioutil.ReadAll(pongResp.Body)
		if err != nil {
			err = fmt.Errorf("error reading pong response body: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if pongResp.StatusCode != http.StatusOK {
			err = fmt.Errorf("invalid pong response code: %d", pongResp.StatusCode)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), pongResp.StatusCode)
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(pongBody)))

		_, err = w.Write(pongBody)
		if err != nil {
			err = fmt.Errorf("error writing ping response: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Pong(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Print(pong)

		defer func() {
			err := r.Body.Close()
			if err != nil {
				logger.Printf("error closing pong request body: %v", err)
			}
		}()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = fmt.Errorf("error reading pong request body: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body = append(body, []byte("\n"+pong)...)

		w.Header().Set("Content-Length", strconv.Itoa(len(body)))

		_, err = w.Write(body)
		if err != nil {
			err = fmt.Errorf("error writing pong response: %w", err)
			logger.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
