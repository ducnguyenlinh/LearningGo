package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type cicle struct {
	radius float64
}

// Circle
func (c cicle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func (c cicle) perim() float64  {
	return 2 * math.Pi * c.radius
}

// Rectangle
func (r rect) area() float64  {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

//
func measure(g geometry)  {
	fmt.Println("geometry: ", g)

	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main()  {
	r := rect{6, 8}
	c := cicle{5}

	measure(r)
	measure(c)
}
