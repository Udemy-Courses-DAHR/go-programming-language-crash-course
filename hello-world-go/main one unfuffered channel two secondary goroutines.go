/* package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel - pure hand-to-hand communication
	taskChannel := make(chan string)

	// GOROUTINE 1: The Task Producer
	go func() {
		tasks := []string{"task A", "task B", "task C", "task D"}

		for _, task := range tasks {
			fmt.Printf("Producer: I have %s ready to hand off\n", task)
			taskChannel <- task // Producer BLOCKS here until Consumer takes it
			fmt.Printf("Producer: %s was taken, preparing next...\n", task)
			time.Sleep(1000 * time.Millisecond) // Simulate work
		}

		close(taskChannel) // No more tasks
	}()

	// GOROUTINE 2: The Task Consumer
	go func() {
		for task := range taskChannel {
			fmt.Printf("Consumer: I received %s, now processing...\n", task)
			time.Sleep(3 * time.Second) // Simulate processing time
			fmt.Printf("Consumer: Finished processing %s\n", task)
		}
		fmt.Println("Consumer: All tasks completed!")
	}()

	// Main goroutine just waits for both to finish
	time.Sleep(20 * time.Second)
	fmt.Println("Main: Program complete")
}
*/