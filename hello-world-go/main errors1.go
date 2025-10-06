/* package main

import (
	"errors" // Import the errors package
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// run returns an error if the input number is 0.
// Otherwise, it returns nil.
func run(number int) error {
	if number == 0 {
		// Return a custom MyError for a specific failure case.
		return &MyError{
			time.Now(),
			"the input number was 0",
		}
	} else if number < 0 {
		// Return a simple error using errors.New for another failure case.
		return errors.New("the input number was negative")
	}

	// Return nil to indicate success for all other numbers.
	return nil
}

func main() {
	// Case 1: This will trigger an error.
	fmt.Println("--- Test Case 1: Input is 0 (triggers MyError) ---")
	if err := run(0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success: No error occurred.")
	}

	fmt.Println("\n--- Test Case 2: Input is 5 (no error) ---")
	// Case 2: This will not trigger an error.
	if err := run(5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success: No error occurred.")
	}

	fmt.Println("\n--- Test Case 3: Input is -1 (triggers simple error) ---")
	// Case 3: This will trigger a different, simple error.
	if err := run(-1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success: No error occurred.")
	}
}
*/