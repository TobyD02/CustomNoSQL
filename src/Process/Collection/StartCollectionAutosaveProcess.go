package Collection

import (
	"log"
	"time"
)

func StartCollectionAutosaveProcess(interval time.Duration, stop <-chan struct{}) {
	collections, err := GetAllCollectionsSingletonProcess()

	if err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(interval) // start timer
		defer ticker.Stop()                // defer timer stop
		for {
			select { // Wait for an event
			case <-ticker.C: // on ticker interval, save the collections
				err = SaveCollectionsProcess(*collections)
				if err != nil {
					log.Printf("failed to save collections: %v", err)
				}
			case <-stop: // on stop condition - return
				return
			}
		}
	}()
}
