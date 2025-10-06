package car

import "fmt"

type Piston struct{}

func (p Piston) Move() {
	fmt.Println("Piston is moving")
}

type Cylinder struct {
	piston Piston
}

func (c Cylinder) Operate() {
	c.piston.Move()
	fmt.Println("Cylinder is operating")
}

// Composition: Car "has a" Engine and "has" Wheels
type Engine struct {
	cylinder Cylinder
}

func (e Engine) Start() {
	e.cylinder.Operate()
	fmt.Println("Engine started")
}

type Wheels struct{}

func (w Wheels) Roll() {
	fmt.Println("Wheels are rolling")
}

// Car is composed of an Engine and Wheels
type Car struct {
	engine Engine
	wheels Wheels
}

// Car delegates to its components
func (c Car) Drive() {
	c.engine.Start()
	c.wheels.Roll()
	fmt.Println("Car is driving!")
}
