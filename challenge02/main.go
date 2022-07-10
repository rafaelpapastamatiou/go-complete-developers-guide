package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func main() {
	t := Triangle{
		height: 10,
		base:   8,
	}

	s := Square{
		sideLength: 5,
	}

	fmt.Println("Triangle's area: ")
	printArea(t)

	fmt.Println()

	fmt.Println("Square's area: ")
	printArea(s)
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
