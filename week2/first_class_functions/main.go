package main

import "fmt"
import "math"

var funcVar func(int) int

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}	

func getFunc(x int) func(int) int {
	if x == 1 {
		return incFn
	} else {
		return decFn
	}
}

func MakeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	return func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x - oX, 2) + math.Pow(y - oY, 2))
	}
}

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	// Deferred function
	b := 10
	defer fmt.Printf("Deferred call %d\n", b)

	b = 20
	fmt.Printf("Regular call %d\n", b)

	// Function as a variable
	funcVar = incFn
	fmt.Println(funcVar(1)) // 2

	// Function as a parameter
	fmt.Println(applyIt(incFn, 2)) // 3
	fmt.Println(applyIt(decFn, 2)) // 1

	// Anonymous function
	fmt.Println(applyIt(func(x int) int { return x * x }, 3)) // 9

	// Function as a return value
	fmt.Println(applyIt(getFunc(1), 2)) // 3
	fmt.Println(applyIt(getFunc(2), 2)) // 1

	// Closure
	origin := MakeDistOrigin(0, 0)
	fmt.Println(origin(2, 2)) // 2.8284271247461903
	fmt.Println(origin(3, 3)) // 4.242640687119285
	fmt.Println(origin(3, 4)) // 5

	// Variadic function
	fmt.Println(getMax(1, 3, 2, 4, 5)) // 5

}

