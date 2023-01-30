package main

import (
	"fmt"
	"net/http"
)

func proxyServer(write http.ResponseWriter, request *http.Request) {
	prepareRequest(request)
}

func prepareRequest(request *http.Request) {
	//conf := getConfig()
	queryUrl := request.Header.Get("url")
	fmt.Print(queryUrl)
}
