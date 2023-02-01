package main

import (
	"net/http"
)

type RequestCache struct {
	Response http.Response
	Error    error
}

func setCache(write http.ResponseWriter, params *http.Request) {
	//conf := getConfig()
	//if conf.Cache.sign == true {
	//
	//}
	sendProxyResponse(write, sendRequest(params))
}
