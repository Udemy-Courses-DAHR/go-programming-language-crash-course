/* package main

import (
	"fmt"
	"time"
)

func main() {
	runway1 := make(chan string) // Runway for domestic flights
	runway2 := make(chan string) // Runway for international flights

	go func() {
		time.Sleep(100 * time.Millisecond)
		runway1 <- "Flight 123: Domestic flight requesting landing"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		runway2 <- "Flight 456: International flight requesting landing"
	}()

	// The Air Traffic Controller (select) monitors both runways
	select {
	case message := <-runway1:
		fmt.Println("Controller:", message)
	case message := <-runway2:
		fmt.Println("Controller:", message) // This one happens first!
	}
}

// Output: "Controller: Flight 456: International flight requesting landing"
/*  */ 