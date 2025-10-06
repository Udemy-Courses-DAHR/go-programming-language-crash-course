/* package main

import (
	"fmt"
	"time"
)

// 1. Unbuffered Channel (Synchronous)
// Sends block until someone receives. Like a direct hand-off.
func main() {
	ch := make(chan string) // Unbuffered channel

	go func() {
		ch <- "ping" // Send blocks until main is ready to receive.
	}()

	time.Sleep(2 * time.Second) // Ensure the goroutine runs first
	msg := <-ch                 // Receive blocks until the goroutine sends.
	fmt.Println(msg)            // Output: ping
}
*/