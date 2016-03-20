package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type circle struct {
	x, y, r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type rectangle struct {
	x1, y1, x2, y2 float64
}

func (r rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

type shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()

	}
	return area
}

func totalPerimeter(shapes ...shape) float64 {
	var p float64
	for _, s := range shapes {
		p += s.perimeter()

	}
	return p
}

func main() {
	c := circle{0, 0, 5}
	fmt.Println(totalArea(c))
	fmt.Println(totalPerimeter(c))
}
