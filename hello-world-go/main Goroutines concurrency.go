/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 3. Launching multiple workers
func main() {
	// Print a line on console every second with the number of the second since start
	go func() {
		for i := 0; i < 22; i++ {
			fmt.Printf("Second %d  +++++++++++++++++++++++\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// Give the above goroutine a bit of time to start so the tick is printed before the workers finish.
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		go worker(i) // Launches 5 concurrent workers
	}
	time.Sleep(21 * time.Second)
}
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Random sleep to simulate work for id. From 1 second to 20 seconds
	sleepDuration := time.Duration(rand.Intn(20)+1) * time.Second
	fmt.Printf("Worker %d sleeping for %v\n", id, sleepDuration)
	time.Sleep(sleepDuration)
	fmt.Printf("Worker %d done after %v\n", id, sleepDuration)
}
*/