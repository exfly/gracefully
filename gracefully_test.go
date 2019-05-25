package gracefully

import (
	"log"
	"testing"
	"time"
)

func TestGracefullExample(t *testing.T) {
	var stopCh = SetupSignalHandler()
	go func() {
		for {
			log.Println("alive...")
			time.Sleep(1 * time.Second)
		}
	}()
	<-stopCh
	log.Println("gracefully shutdown")
}
