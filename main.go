package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elazarl/goproxy"
	"github.com/joho/godotenv"
)

func main() {
 	godotenv.Load()

	proxy := goproxy.NewProxyHttpServer()

	proxy.OnRequest().DoFunc(func (req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			var headers map[string]string

			headersEnv := os.Getenv("HEADERS")

			err := json.Unmarshal([]byte(headersEnv), &headers)

			if err != nil {
				log.Printf("Error unmarshalling JSON: %v\n", err)
				return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v\n", err))
			}

			for key, value := range headers {
				req.Header.Set(key, value)
			}

			return req, nil
		})

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "3000"
	}
	port, err := strconv.Atoi(portEnv)

	if err != nil {
		log.Fatalf("Invalid port number: %s\n", portEnv)
	}

	verboseEnv := os.Getenv("VERBOSE")
	if verboseEnv == "" {
		verboseEnv = "f"
	}
	verbose, err := strconv.ParseBool(verboseEnv)

	if err != nil {
		log.Fatalf("Invalid verbose value: %s\n", verboseEnv)
	}

	proxy.Verbose = verbose

	log.Printf("Proxy server is running on http://%s:%d", host, port)
	listenErr := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), proxy)

	if listenErr != nil {
		log.Fatalf("Failed to start server: %v", listenErr)
	}
}
