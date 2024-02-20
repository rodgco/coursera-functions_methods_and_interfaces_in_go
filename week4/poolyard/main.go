package main

import "fmt"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Point struct {
	x, y float64
}

type Triangle struct {
	a, b, c Point
}

func (t Triangle) Area() float64 {
	return 0.5 * ((t.a.x-t.c.x)*(t.b.y-t.a.y) - (t.a.x-t.b.x)*(t.c.y-t.a.y))
}

func (t Triangle) Perimeter() float64 {
	return distance(t.a, t.b) + distance(t.b, t.c) + distance(t.c, t.a)
}

type Circle struct {
	center Point
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Rectangle struct {
	a, b Point
}

func (r Rectangle) Area() float64 {
	return distance(r.a, r.b) * distance(r.a, r.b)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (distance(r.a, r.b) + distance(r.b, r.a))
}

func distance(a, b Point) float64 {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}

func fitInYard(s Shape2D) bool {
	return s.Area() <= 100 && s.Perimeter() <= 50
}

func main() {
	t := Triangle{Point{0, 0}, Point{3, 0}, Point{0, 4}}
	c := Circle{Point{0, 0}, 5}
	r := Rectangle{Point{0, 0}, Point{3, 4}}

	shapes := []Shape2D{t, c, r}

	for _, s := range shapes {
		if fitInYard(s) {
			fmt.Printf("%T fits in the yard\n", s)
		} else {
			fmt.Printf("%T does not fit in the yard\n", s)
		}
	}
}
