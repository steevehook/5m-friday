package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type calculator struct {
}

func (a calculator) areaSum(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).length
			sum += l * l
		}
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s))
}
