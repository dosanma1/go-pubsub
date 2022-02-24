package main

import (
	"log"
	"sync"
	"time"

	ps "github.com/dosanma1/go-pubsub"
)

func main() {
	var wg sync.WaitGroup

	ps := ps.NewPubsub()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(<-ps.Subscribe("AAPL"))
	}()

	time.Sleep(50 * time.Millisecond)

	ps.Publish("AAPL", 160.07)

	wg.Wait()
}
