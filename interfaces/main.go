package main

import (
	"fmt"
	"math"
)

// defining an interface when you want to have a common behaviour, need flexible function parameters, or for testing purposes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

// methods matching the interface's methods signature
// Circle type implements the shape interface because it has 'Area' and 'Parameter' methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle type implements the shape interface because it has 'Area' and 'Parameter' methods
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// I am general function that serves any object which has (a) function(s) that matches the interface template.
// I allow polymorphism by enabling different types to be used interchangeably.
func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println("--")
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	PrintShapeDetails(c)
	PrintShapeDetails(r)
}
