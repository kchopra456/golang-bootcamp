package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

func (p *Circle) area() float64 {
	return math.Pi * p.radius * p.radius
}

func (p *Square) area() float64 {
	return p.length * p.length
}

type Shape interface {
	area() float64
}

func sumAreas() {
	shapes := []Shape{&Square{10}, &Circle{10}}

	sum := 0.0

	for _, sh := range shapes {
		sum += sh.area()
	}
	fmt.Printf("Cumulative area sum: %f", sum)
}
