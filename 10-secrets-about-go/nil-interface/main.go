package main

import "fmt"

type person struct {
}

func main() {
	p := adult(10)
	if p != nil {
		fmt.Println(":(, I'm still a kid")
	} else {
		fmt.Println("He-he, finally got adult")
	}
}

func adult(n int) interface{} {
	var res *person = nil
	if n < 18  {
		return nil // type: *person | value: nil
	}
	return res // type: nil | value: nil
}
