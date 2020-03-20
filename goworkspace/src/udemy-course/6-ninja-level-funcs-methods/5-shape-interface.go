package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea()
}

type Square struct {
	length float64
	width  float64
}

func (s Square) getArea() float64 {
	return s.length * s.width
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	pi := 3.14
	area := math.Pow(c.radius, 2) * pi
	return area
}

func main() {
	mySquare := Square{
		length: 25,
		width:  120,
	}

	fmt.Printf("mySquare area = %v\n", mySquare.getArea())

	myCircle := Circle{radius: 2}

	fmt.Printf("myCircle area = %v\n", myCircle.getArea())
}
