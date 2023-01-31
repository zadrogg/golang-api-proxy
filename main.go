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
	// получаем переменную сервера
	server := os.Getenv("SERVER")
	// определяем какие роуты будет принимать прокси
	http.HandleFunc("/", proxyServer)
	// запускаем сервер и пишем логи о ошибках
	errorHandler(http.ListenAndServe(server, nil))
}

func errorHandler(error error) {
	if error != nil {
		_, _ = os.Stderr.WriteString(error.Error())
	}
}
