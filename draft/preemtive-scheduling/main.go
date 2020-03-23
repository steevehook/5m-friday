package main

import "fmt"

func main() {
	done := false

	go func(){
		fmt.Println("Here")
		done = true
	}()

	for !done {
	}
	fmt.Println("done!")
}
