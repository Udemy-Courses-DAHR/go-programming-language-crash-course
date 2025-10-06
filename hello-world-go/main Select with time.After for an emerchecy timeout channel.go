/* package main

import (
	"fmt"
	"time"
)

func main() {
	landingRequest := make(chan string)
	timeout := time.After(2 * time.Second) // Emergency timeout channel

	go func() {
		// Simulate a plane taking too long to respond
		time.Sleep(3 * time.Second)
		landingRequest <- "Flight 789: Finally ready to land"
	}()

	// Controller must respond to either the plane OR the timeout
	select {
	case request := <-landingRequest:
		fmt.Println("Approved:", request)
	case <-timeout:
		fmt.Println("EMERGENCY: Plane unresponsive - alert security!") // This happens!
	}
}

// Output: "EMERGENCY: Plane unresponsive - alert security!"
*/