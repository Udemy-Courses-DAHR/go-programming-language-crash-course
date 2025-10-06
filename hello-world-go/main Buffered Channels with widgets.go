/* package main

import (
	"fmt"
	"time"
)

func main() {
	// A conveyor belt that can hold 2 widgets
	conveyor := make(chan string, 2)

	// Producer: Makes widgets
	go func() {
		widgets := []string{"widgetA", "widgetB", "widgetC", "widgetD", "widgetE"}

		for _, widget := range widgets {
			if len(conveyor) == cap(conveyor) {
				fmt.Printf("Producer: Belt FULL with %d items, waiting...\n", len(conveyor))
			}

			// Current status of the belt
			fmt.Printf("Producer: Belt has %d/%d items\n",
				len(conveyor), cap(conveyor))

			fmt.Printf("Producer: Making %s\n", widget)
			time.Sleep(2000 * time.Millisecond) // Production time

			conveyor <- widget // Put on conveyor belt

			fmt.Printf("Producer: %s placed on belt. Belt now has %d/%d items\n",
				widget, len(conveyor), cap(conveyor))
		}
		close(conveyor)
	}()

	// Consumer: Packages widgets (slower than producer) And starts later

	go func() {
		time.Sleep(3 * time.Second) // Start after a delay of at least one item on belt to avoid race condition
		for widget := range conveyor {
			fmt.Printf("Consumer: Belt has %d items, taking %s for packaging\n",
				len(conveyor), widget)
			time.Sleep(4 * time.Second) // Packaging takes longer
			fmt.Printf("Consumer: %s packaged! Belt now has %d items\n",
				widget, len(conveyor))
		}
	}()

	time.Sleep(24 * time.Second)
}
*/