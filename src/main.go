package main

import (
	"context"
	"github.com/rimonmostafiz/go-microservice-playgroud/src/handlers"
	log2 "log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log2.New(os.Stdout, "go-micro", log2.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	S := &http.Server{
		Addr:              ":9090",
		Handler:           serveMux,
		TLSConfig:         nil,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	go func() {
		err := S.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	logger.Println("Received terminate, graceful shutdown, signal type: ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	S.Shutdown(tc)
}
