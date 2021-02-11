package main

import (
	"http-mux/internal/http_server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	httpServer := http_server.New()
	go func() {
		if err := httpServer.Start(); err != nil {
			// Couldn't start. Do os.Exit(1)
			shutdown <- syscall.SIGUNUSED
		}
	}()

	s := <-shutdown
	if s == syscall.SIGUNUSED {
		log.Println("failure shutdown")
		os.Exit(1)
	}

	log.Println("graceful shutdown...")
	httpServer.Stop()
	log.Println("bye")
}
