package main

import (
	"fmt"
	"strings"
)

func main() {
	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	people := []string{
		"John",
		"Steve",
		"Jane",
	}

	//p := people[1:2]
	p := people[2:3:3]
	fmt.Println("capacity:", cap(p))
	p = append(p, "Mike")
	inspect("people: ", people)
	inspect("sub people: ", p)
	fmt.Println(people[2])

	peopleClone := append(people[:0:0], people...)
	inspect("people: ", people)
	inspect("peopleClone: ", peopleClone)
}

func inspect(label string, people []string) {
	f := strings.Repeat("%p | ", len(people))
	f = label + f + "\n"
	var args []interface{}
	for i := range people {
		args = append(args, &people[i])
	}
	fmt.Printf(f, args...)
}
