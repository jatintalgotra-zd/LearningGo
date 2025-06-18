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
