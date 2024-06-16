package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) getX() float64 {
	return p.x
}
func (p *Point) getY() float64 {
	return p.y
}
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {
	A := NewPoint(12, 13)
	B := NewPoint(20, 65)
	fmt.Println(Disctance(A, B))

}

func Disctance(a *Point, b *Point) float64 {
	return math.Sqrt((a.getX()-b.getX())*(a.getX()-b.getX()) + (a.getY()-b.getY())*(a.getY()-b.getY()))
}
