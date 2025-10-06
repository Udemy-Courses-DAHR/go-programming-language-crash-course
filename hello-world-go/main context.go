// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func longOperation(ctx context.Context) {
// 	for i := 0; i < 5; i++ {
// 		select {
// 		case <-time.After(1 * time.Second):
// 			fmt.Println("Doing work...", i)
// 		case <-ctx.Done():
// 			// The context was cancelled, stop working.
// 			fmt.Println("Operation cancelled:", ctx.Err())
// 			return
// 		}
// 	}
// 	fmt.Println("Operation completed successfully!")
// }

// func main() {
// 	// Create a context that cancels after 2.5 seconds
// 	ctx, cancel := context.WithTimeout(context.Background(), 5500*time.Millisecond)
// 	defer cancel() // Good practice: release resources if operation completes before timeout

// 	go longOperation(ctx)

// 	// Wait a bit to see the output (in a real program, you'd use a WaitGroup)
// 	time.Sleep(10 * time.Second)
// 	fmt.Println("Main function completed")
// }
