package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	sq := square{
		length: 2,
	}

	cr := circle{
		radius: 2,
	}

	info(sq)
	info(cr)
}

func (s square) area() float64 {
	areaOfSquare := s.length * s.length
	return areaOfSquare
}

func (c circle) area() float64 {
	areaOfCircle := (math.Pi) * (c.radius * c.radius)
	return areaOfCircle
}

func info(s shape) {
	fmt.Println(s.area())
}
