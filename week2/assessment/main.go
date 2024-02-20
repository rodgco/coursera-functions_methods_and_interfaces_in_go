package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}


func main() {
	fmt.Print("Please inform the acceleration: ")
	var acceleration float64
	fmt.Scan(&acceleration)

	fmt.Print("Please inform the initial velocity: ")
	var initialVelocity float64
	fmt.Scan(&initialVelocity)

	fmt.Print("Please inform the initial displacement: ")
	var initialDisplacement float64
	fmt.Scan(&initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Print("\nPlease inform the time: ")
	var time float64
	fmt.Scan(&time)

	fmt.Println("\nThe displacement is: ", fn(time))
}

