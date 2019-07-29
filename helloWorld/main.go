package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius*c.radius*math.Pi
}

type Shape interface {
	area() float64
}

func info (z Shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Circle{10}
	info(s)
}
 