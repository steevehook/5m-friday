package main

import "fmt"

func main() {
	s := "hello world!"
	for i, v := range s {
		fmt.Println(i, v)
	}
}
