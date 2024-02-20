package main

import "fmt"
import "rodgco/coursera-functions_methods_and_interfaces_in_go/week3/classes/data" 

type Point struct {
	X, Y int
}

type Triangle struct {
	a, b, c Point
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func (t *Triangle) Move(dx, dy int) {
	t.a.Move(dx, dy)
	t.b.Move(dx, dy)
	t.c.Move(dx, dy)
}

func main() {
	t := Triangle{Point{1, 2}, Point{3, 4}, Point{5, 6}}
	t.Move(10, 10)

	fmt.Println(t)

	d := data.New("A Data", 10)

	fmt.Println(d.Name())
	fmt.Println(d.Value())
	d.Double()
	fmt.Println(d.Value())
}	
