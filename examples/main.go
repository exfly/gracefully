package main

import (
	"log"
	"time"

	"github.com/exfly/gracefully"
)

func main() {
	var stopCh = gracefully.SetupSignalHandler()
	go func() {
		for {
			log.Println("alive...")
			time.Sleep(1 * time.Second)
		}
	}()
	<-stopCh
	log.Println("gracefully shutdown")
}
