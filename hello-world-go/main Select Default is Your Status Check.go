/* package main

import (
	"fmt"
	"time"
)

func main() {
	emergencyChannel := make(chan string)
	routineChannel := make(chan string)

	// No emergencies right now - channels are empty
	select {
	case emergency := <-emergencyChannel:
		fmt.Println("URGENT:", emergency)
	case routine := <-routineChannel:
		fmt.Println("Routine:", routine)
	default:
		fmt.Println("All clear - monitoring frequencies") // This runs immediately!
	}

	// Now let's simulate an emergency coming in
	go func() {
		emergencyChannel <- "Mayday! Engine failure!"
	}()

	time.Sleep(100 * time.Millisecond) // Give the emergency time to come in

	select {
	case emergency := <-emergencyChannel:
		fmt.Println("URGENT:", emergency) // This runs now!
	default:
		fmt.Println("All clear - monitoring frequencies")
	}
}

// Output:
// "All clear - monitoring frequencies"
// "URGENT: Mayday! Engine failure!"
*/