/* package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this Goroutine is done, no matter how it exits
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Add 1 to the counter for each Goroutine we start
		go worker(i, &wg)
	}

	wg.Wait() // Block here until all Goroutines call wg.Done()
	fmt.Println("All workers completed")
}
*/