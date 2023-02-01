package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type UserRequest struct {
	url string
}

func init() {
	//load values from .env
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	conf := getConfig()
	// открываем соединение с хранением кэша (storage Example: redis)
	//if conf.Cache.sign == true {
	//	storage := getConnection
	//	defer storage.Close()
	//}
	// определяем какие роуты будет принимать прокси
	http.HandleFunc("/", proxyServer)
	// запускаем сервер и пишем логи о ошибках
	errorHandler(http.ListenAndServe(conf.Server.url, nil))
}

func errorHandler(error error) {
	if error != nil {
		_, _ = os.Stderr.WriteString(error.Error())
	}
}
