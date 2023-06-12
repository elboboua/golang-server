package main

import (
	"errors"
	"example/server/api"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Seed random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Set up server
	router := api.NewRouter()
	api.AddRoutingGroups(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Started example server")
		}
	}()

	// Handling a graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit 
}