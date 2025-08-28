package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(p2 Point) float64 {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	fmt.Printf("Точка 1: (%.2f, %.2f)\n", p.x, p.y)
	fmt.Printf("Точка 2: (%.2f, %.2f)\n", p2.x, p2.y)

	distance := p.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
