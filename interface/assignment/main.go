package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	sq := square{sideLength: 10}
	tr := triangle{base: 5, height: 23}

	printArea(sq)
	printArea(tr)
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
