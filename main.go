package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"./artists"
	"./cache"
	"./io/postgres"
)

func cleanup() {
	fmt.Println("Shutting down api...")
	postgres.Close()
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
	postgres.Initialize()
	cache.Initialize()
	fmt.Println("Starting...")
	http.Handle("/artists/", artists.Handler())
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":8080", nil))
}
