package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		req.Header.Add("X-Custom-Header", "Value")

		return req, nil
	})

	log.Println("Proxy server is running on http://localhost:3000")
	err := http.ListenAndServe(":3000", proxy)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
