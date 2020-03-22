package main

import "fmt"

func main() {
	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	people := []string{
		"John",
		"Steve",
		"Jane",
	}
	fmt.Printf("address of people[1]: %p\n", &people[1])

	oneCap3 := people[1:2]
	fmt.Println("capacity of oneCap3:", cap(oneCap3))
	fmt.Printf("address of oneCap3[0]: %p\n", &oneCap3[0])

	oneCap3 = append(oneCap3, "Alex")
	fmt.Println(people[2]) // changed
	fmt.Printf("address of oneCap3[1]: %p\n", &oneCap3[1])
	fmt.Printf("address of people[2]: %p\n", &people[2])

	oneCap1 := people[2:3:3]
	fmt.Println("capacity of oneCap1:", cap(oneCap1))

	peopleClone := append(people[:0:0], people...)
	fmt.Printf("original : %p | %p | %p\n", &people[0], &people[1], &people[2])
	fmt.Printf("clone : %p | %p | %p\n", &peopleClone[0], &peopleClone[1], &peopleClone[2])

	peopleClone[0] = "changed"
	fmt.Println(people[0]) // did not change
}
