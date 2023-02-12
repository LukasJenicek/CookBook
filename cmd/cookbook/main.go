package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LukasJenicek/CookBook/services/api"
)

func main() {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	srv := http.Server{
		Addr:    ":3000",
		Handler: api.Routes(),
	}

	go func() {
		select {
		case sig := <-signals:
			fmt.Printf("Got %s signal. Shutting down the server...\n", sig)
			srv.Close()
		}
	}()

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
