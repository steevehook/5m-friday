package main

import (
	"fmt"
)

type person struct {
	name string
}

func info(p *person) {
	fmt.Println(p.name)
}

func main() {
	steve := new(person)
	steve.name = "steve"
	john := &person{"john"}
	info(steve)
	info(john)
}
