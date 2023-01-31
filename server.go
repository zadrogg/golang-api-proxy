package main

import (
	"io"
	"net/http"
	"net/url"
)

func proxyServer(write http.ResponseWriter, request *http.Request) {
	params, err := prepareRequest(paramRequest(request), request)
	if err != nil {
		return
	}

	setCache()

	sendProxyResponse(write, sendRequest(params))
}

func paramRequest(request *http.Request) string {
	queryUrl, _ := url.Parse(request.RequestURI)
	values := queryUrl.Query()

	return values.Get("url")
}

func prepareRequest(url string, request *http.Request) (*http.Request, error) {
	return http.NewRequest("POST", url, request.Body)
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

	write.WriteHeader(resp.StatusCode)
	_, _ = write.Write(bodyBytes)
}
