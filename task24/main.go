package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

//Конструктор для Point
func NewPoint() *Point {
	p := new(Point)
	//Значения по умолчанию
	p.x = 3
	p.y = 4
	return p
}
func (p Point) distance() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}
func main() {
	pnt := NewPoint()
	fmt.Println(pnt.distance())
}
