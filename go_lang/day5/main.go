package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func Add(a int, b int) int {
	return a + b
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func Info(s shape) float64 {
	return s.Area()
}

func main() {

	fmt.Println("Add func returns", Add(6, 5))

	s1 := Square{
		length: 3,
		width:  3,
	}
	fmt.Printf("Area of square is %f\n", s1.Area())

	c1 := Circle{
		radius: 3,
	}
	fmt.Printf("Area of circle is %f\n", c1.Area())

	fmt.Println("Info Square area", Info(s1))
	fmt.Println("Info Circle area", Info(c1))
}
