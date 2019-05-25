package main

import (
	"log"
	"testing"
	"time"

	"github.com/exfly/gracefully"
)

func TestGracefullExample(t *testing.T) {
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
