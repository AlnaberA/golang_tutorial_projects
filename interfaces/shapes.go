package main

import (
	"fmt"
	"math"
)

type geomtry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

//Method that operates on rectangle
//(rect rectangle) is the object we want the method to operate on.
func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

//Method that operates on circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geomtry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	rectangle := rectangle{width: 3, height: 4}
	measure(rectangle)
}
