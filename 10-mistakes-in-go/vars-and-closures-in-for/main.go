package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []int{5, 10, 15} {
		// go routines must be scheduled first
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)
}
