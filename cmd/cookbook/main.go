package main

import (
	"log"
	"net/http"

	"github.com/LukasJenicek/CookBook/services/api"
)

func main() {
	err := http.ListenAndServe(":3000", api.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
