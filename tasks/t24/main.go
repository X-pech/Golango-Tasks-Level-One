package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// This is what I learnt the conception of OOP from

func NewPoint(x float64, y float64) Point {
	p := new(Point)
	p.x = x
	p.y = y
	return *p
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func Distance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func main() {
	p1 := NewPoint(0.0, 3.0)
	p2 := NewPoint(4.0, 0.0)
	fmt.Println(Distance(p1, p2))
}
