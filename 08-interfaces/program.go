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

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * r.Height * 2 * r.Width
}

type ShapeWithArea interface {
	Area() float32
}

func printArea(shape ShapeWithArea) {
	fmt.Println("Area = ", shape.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(shape ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shape.Perimeter())
}

/*
type Dimension interface {
	Area() float32
	Perimeter() float32
}
*/

//interface composition
type Dimension interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printDimension(shape Dimension) {
	printArea(shape)
	printPerimeter(shape)
}

/*
	Perimeter for Circle = 2 * pi * r
	Perimeter for Rectangle = 2 * Height + 2 * Width
*/

func main() {
	c := &Circle{Radius: 12}
	/* printArea(c)
	printPerimeter(c) */
	printDimension(c)

	r := &Rectangle{Height: 8, Width: 12}
	/* printArea(r)
	printPerimeter(r) */
	printDimension(r)
}
