package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

//Конструктор для Point
func NewPoint(x, y int) *Point {
	p := new(Point)
	//Значения по умолчанию
	p.x = x
	p.y = y
	return p
}
func (p Point) distance(t *Point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-t.x), 2) + math.Pow(float64(p.y-t.y), 2))
}
func main() {
	pnt1 := NewPoint(0, 0)
	pnt2 := NewPoint(3, 4)

	fmt.Println(pnt1.distance(pnt2))
}
