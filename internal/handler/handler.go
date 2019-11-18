package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	ping           = "Ping!"
	pong           = "Pong!"
	pongServiceURL = "http://192.168.99.100:32724"
)

func Ping(logger *log.Logger, httpClient *http.Client) http.HandlerFunc {
	logger.Print(ping)

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := r.Body.Close()
			if err != nil {
				logger.Printf("Error closing request body: %v", err)
			}
		}()

		_, err := io.Copy(ioutil.Discard, r.Body)
		if err != nil {
			logger.Printf("Error discarding request body: %v", err)
			return
		}

		buf := bytes.NewBufferString(ping)

		pongReq, err := http.NewRequest(http.MethodPost, pongServiceURL, buf)
		if err != nil {
			logger.Printf("Error creating pong request body: %v", err)
			return
		}

		pongResp, err := httpClient.Do(pongReq)
		if err != nil {
			logger.Printf("Error performing pong request body: %v", err)
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
			logger.Printf("Error reading pong response body: %v", err)
			return
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(pongBody)))

		_, err = w.Write(pongBody)
		if err != nil {
			logger.Printf("Error writing ping response body: %v", err)
		}
	}
}

func Pong(logger *log.Logger) http.HandlerFunc {
	logger.Print("Pong!")

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := r.Body.Close()
			if err != nil {
				logger.Printf("Error closing request body: %v", err)
			}
		}()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Printf("Error reading request body: %v", err)
			return
		}

		body = append(body, []byte("\n"+pong)...)

		w.Header().Set("Content-Length", strconv.Itoa(len(body)))

		_, err = w.Write(body)
		if err != nil {
			logger.Printf("Error writing response body: %v", err)
		}
	}
}
