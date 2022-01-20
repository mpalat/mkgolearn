package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

type ShapeWithArea interface {
	Area() float64
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

//interface composition
type ShapeWithDetails interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func main() {

	c := Circle{Radius: 12}
	// fmt.Println(c.Area())
	// PrintArea(c)
	// PrintPerimeter(c)
	PrintDetails(c)

	r := Rectangle{height: 10, width: 12}
	// fmt.Println(r.Area())
	// PrintArea(r)
	// PrintPerimeter(r)
	PrintDetails(r)

}

func PrintDetails(x ShapeWithDetails) {
	// Print Area and Perimeter
	PrintArea(x)
	PrintPerimeter(x)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func PrintArea(x ShapeWithArea) {
	fmt.Println("Area:", x.Area())
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}
func PrintPerimeter(x ShapeWithPerimeter) {
	fmt.Println("Perimenter:", x.Perimeter())
}
