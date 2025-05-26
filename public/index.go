package main

import (
	"CustomNoSQL/src/Process/Collection"
	"CustomNoSQL/src/Route"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	Route.RegisterRoutes(r)
	// Load singletons
	collections, err := Collection.GetAllCollectionsSingletonProcess()

	if err != nil {
		panic(err)
	}

	stop := make(chan struct{})
	Collection.StartCollectionAutosaveProcess(5*time.Second, stop)

	// Handle process termination
	signalChannel := make(chan os.Signal, 1)                    // Create channel to receive OS signals
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) // Register the channel

	go func() { // Create goroutine that blocks until the signal is received
		<-signalChannel
		err = Collection.SaveCollectionsProcess(*collections)
		if err != nil {
			panic(err)
		}

		close(stop)
		os.Exit(0)
	}()

	r.Run()
}
