package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}

type ShapeWithArea interface {
	Area() float32
}

func printArea(shape ShapeWithArea) {
	fmt.Println("Area = ", shape.Area())
}

func main() {
	c := &Circle{Radius: 12}
	printArea(c)

	r := &Rectangle{Height: 8, Width: 12}
	printArea(r)
}
