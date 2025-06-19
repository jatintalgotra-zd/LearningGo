package main

import (
	"fmt"
	"math"
)

// Shape
// interface which defines Area() method returning float64 value
type Shape interface {
	Area() float64
}

// Rectangle
// struct defining rectangle's length and width
type Rectangle struct {
	length, width float64
}

// Area
// method for Rectangle to calculate area
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Circle
// struct defining circle's radius
type Circle struct {
	radius float64
}

// Area
// method for Circle to calculate area
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {
	r := Rectangle{12, 34}
	c := Circle{10}

	fmt.Println("Area of Rectangle: ", CalculateArea(r))
	fmt.Println("Area of Circle: ", CalculateArea(c))
}

// CalculateArea
// function which takes a Shape argument and returns its Area
func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

// Cost
// function with argument of type Shape which returns cost based on Area and rates of each type
func Cost(s Shape) float64 {
	// 0.02 for sq, 0.042 for rec, 1.20 for cir
	area := s.Area()
	switch s.(type) {
	case Square:
		rate := 0.02
		return area * rate

	case Rectangle:
		rate := 0.42
		return area * rate

	case Circle:
		rate := 1.20
		return area * rate

	default:
		rate := 1.00
		return area * rate
	}
}
