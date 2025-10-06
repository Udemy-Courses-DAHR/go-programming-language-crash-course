/* package main

import (
	"fmt"
	"hello-world-go/car"
) // Import to run car composition example

type Point struct {
	X, Y int
}

func main() {
	// A regular struct
	p := Point{1, 2}
	fmt.Printf("Just the value (%%v): %v\n", p)
	fmt.Printf("Value + Type (%%T): %T\n", p)
	fmt.Printf("EVERYTHING (%%#v): %#v\n", p) // <- The one you want
	fmt.Println()

	// A pointer to the struct
	ptr := &p
	fmt.Printf("Pointer value (%%v): %v\n", ptr)        // Just the hex address
	fmt.Printf("Pointer type (%%T): %T\n", ptr)         // *main.Point
	fmt.Printf("Pointer EVERYTHING (%%#v): %#v\n", ptr) // (*main.Point)(0x1040a128)
	fmt.Println()

	// A slice
	sl := []int{1, 2, 3}
	fmt.Printf("Slice (%%#v): %#v\n", sl) // []int{1, 2, 3}

	myCar := car.Car{}
	myCar.Drive()
}
*/