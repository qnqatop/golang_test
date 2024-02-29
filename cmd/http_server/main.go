package main

import (
	"fmt"
	"golang_testovoe/cmd/http_server/config/initModules"
	"golang_testovoe/cmd/http_server/config/router"
	"net/http"
	"os"
)

func init() {
	initModules.NewConfig()
}

func main() {
	r := router.Router()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print("Server Started")
}
