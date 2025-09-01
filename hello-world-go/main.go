package main

import (
	"myapp/doctor"
)

func main() {

	var whatToSay string = doctor.Intro()

	println(whatToSay)

}
