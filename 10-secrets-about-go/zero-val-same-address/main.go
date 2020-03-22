package main

import "fmt"

type person struct {
}

func main() {
	p1 := &person{}
	p2 := &person{}
	fmt.Printf("%p\n%p", p1, p2)
}
