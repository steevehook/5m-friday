package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) name() string {
	return "square"
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
	name() string
}

type outputter struct {
}

func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area  of the %s: %f", s.name(), s.area())
}

func (o outputter) JSON(s shape) string {
	res := struct {
		Area float64 `json:"area"`
	}{
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	out := outputter{}
	fmt.Println(out.Text(c))
	fmt.Println(out.JSON(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.JSON(s))
}
