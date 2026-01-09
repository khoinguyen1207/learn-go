package main

import (
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from channel 2"
	}()

	// Wait and receive messages from both channels, 5 seconds apart
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)

	// Using select to wait on multiple channels
	for range 2 {
		select {
		case msg1 := <-ch1:
			println("Received:", msg1)
		case msg2 := <-ch2:
			println("Received:", msg2)
		}
	}
}
