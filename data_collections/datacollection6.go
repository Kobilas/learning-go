package main

import (
	"fmt"
	"math"
)

type Cube struct {
	length float64
}

type Box struct {
	length float64
	width  float64
	height float64
}

type Sphere struct {
	radius float64
}

type Shape interface {
	volume() float64
}

func (c *Cube) volume() float64 {
	return math.Pow(c.length, 3)
}

func (b *Box) volume() float64 {
	return b.length * b.width * b.height
}

func (s *Sphere) volume() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func main() {
	c := Cube{}
	c.length = 3.5
	var sCube Shape
	sCube = &c
	fmt.Println("Volume of cube:", sCube.volume())

	b := Box{}
	b.length = 3.5
	b.width = 4.5
	b.height = 4.2
	sBox = &b
	fmt.Println("Volume of box:", b.volume())
}
