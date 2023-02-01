package main

import (
	"io"
	"log"
	"net/http"
)

func proxyServer(write http.ResponseWriter, request *http.Request) {
	params, err := prepareRequest(request)
	if err != nil {
		return
	}

	setCache(write, params)
}

func prepareRequest(request *http.Request) (*http.Request, error) {
	proxyRequest, err := http.NewRequest("POST", request.Header.Get("Proxy-Url"), request.Body)
	proxyRequest.Header.Set("X-Secret", request.Header.Get("X-Secret"))
	proxyRequest.Header.Set("Authorization", request.Header.Get("Authorization"))

	if err != nil {
		log.Fatal(err)
	}

	return proxyRequest, nil
}

func sendRequest(request *http.Request) *http.Response {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}

func sendProxyResponse(write http.ResponseWriter, resp *http.Response) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(resp.StatusCode)
	_, _ = write.Write(bodyBytes)
}
