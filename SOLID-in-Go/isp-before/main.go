package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

// square is a flat surface -> does not have volume
func (s square) volume() float64 {
	return 0
}

type cube struct {
	length float64
}

func (c cube) area() float64 {
	return math.Pow(c.length, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type shape interface {
	area() float64
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s := square{length: 3}
	c := cube{length: 4}
	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))
}
