package main

import (
	"golang_testovoe/cmd/http_server/config/initModules"
	"golang_testovoe/cmd/http_server/config/router"
	"log"
	"net/http"
	"os"
)

func init() {
}

func main() {
	r := router.Router()
	initModules.NewConfig()

	err := http.ListenAndServe(":8080", r)
	log.Default().Print("Server Started")
	if err != nil {
		os.Exit(1)
	}
}
