package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/devdimidved/int-bst-http/api"
	"github.com/devdimidved/int-bst-http/bst"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var jsonPath = flag.String("json", "./data/input.json", "json file with input integers")
	flag.Parse()

	input := readInput(*jsonPath)

	httpLogger := logrus.New()
	httpLogger.SetFormatter(&logrus.JSONFormatter{})

	var bstSrv bst.Service
	bstSrv = bst.NewService(input)

	bstLogger := logrus.New()
	bstLogger.SetFormatter(&logrus.JSONFormatter{})
	bstSrv = bst.NewLoggingService(bstLogger, bstSrv)

	app := api.NewApplication(httpLogger, bstSrv)

	srv := &http.Server{
		Addr:         *httpAddr,
		Handler:      app,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main : start listening on %s", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("error: listening and serving: %s", err)

	case <-shutdown:
		log.Println("main : start shutdown")

		// Give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown.
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("main : graceful shutdown did not complete in %v : %v", timeout, err)
			err = srv.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
}

func readInput(jsonPath string) []int {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Println("will use an empty array because of can't open file: ", err)
		return []int{}
	}
	defer jsonFile.Close()
	var input []int
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&input)
	if err != nil {
		log.Println("will use an empty array because of can't decode JSON: ", err)
		return []int{}
	}
	return input
}
