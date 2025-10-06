/* package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel for passing ingredients between two chefs
	ingredientChannel := make(chan string)

	// GOROUTINE 1: Prep Chef (chops vegetables)
	go func() {
		vegetables := []string{"carrot", "onion", "potato", "celery"}

		for _, veg := range vegetables {
			fmt.Printf("Prep Chef: Chopping %s...\n", veg)
			time.Sleep(800 * time.Millisecond) // Chopping time
			fmt.Printf("Prep Chef: %s chopped, handing to Cook Chef\n", veg)
			ingredientChannel <- veg // BLOCKS until Cook Chef takes it
			fmt.Printf("Prep Chef: %s was taken, starting next vegetable\n", veg)
		}

		close(ingredientChannel)
		fmt.Println("Prep Chef: All vegetables prepped!")
	}()

	// GOROUTINE 2: Cook Chef (cooks vegetables)
	go func() {
		for vegetable := range ingredientChannel {
			fmt.Printf("Cook Chef: Received %s, now cooking...\n", vegetable)
			time.Sleep(1 * time.Second) // Cooking time
			fmt.Printf("Cook Chef: %s is cooked and ready!\n", vegetable)
		}
		fmt.Println("Cook Chef: All cooking complete!")
	}()

	// Main goroutine is just a bystander
	time.Sleep(10 * time.Second)
}
*/