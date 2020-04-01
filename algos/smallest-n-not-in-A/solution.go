package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{12,0-1,-3,1,2,3,3,3,10}
	fmt.Println(Solution(in))
}

func Solution(T []int) int {
	sort.Ints(T)
	unique := map[int]bool{}
	increment := 1
	for i, element := range T {
		if element > 0 && !unique[element] {
			unique[element] = true
			increment++
		}
		isLast := i+1 > len(T)-1
		if isLast {
			if len(unique) == len(T) {
				return increment
			}
			return 1
		}
		if increment != element && increment < T[i+1] {
			return increment
		}
	}
	return 1
}
