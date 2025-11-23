package main

import (
	"fmt"
)

type Polygon interface {
	Quad
	Shape
}

type Quad interface {
	Area() float64
	isFourSided() bool
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type EmptyShape interface{}

func main() {
	var shape Shape
	squareOne := Square{10}
	fmt.Println(squareOne)
	circleOne := NewCircle(20)
	fmt.Println(circleOne)

	shape = circleOne
	shape = squareOne
	fmt.Println(shape.Area())

	var age EmptyShape
	age = 1.9
	age = "string"
	fmt.Println(age)

}

type Square struct {
	Length float64
}

func NewSquare(length float64) Square {
	return Square{length}
}
func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

func (s Square) String() string {
	return "this is a square"
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}
func (c *Circle) Area() float64 {
	return c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * c.Radius
}
