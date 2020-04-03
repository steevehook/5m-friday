package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sol(6, 1, 1))
}

func sol(A, B, C int) string {
	var res string
	if A+B+C <= 0 {
		return ""
	}
	aTimes, bTimes, cTimes := A/2, B/2, C/2
	for i := 0; i < aTimes+bTimes+cTimes; i++ {
		if A == 1 {
			res += "a"
		}
		if aTimes > 0 {
			res += strings.Repeat("a", 2)
			aTimes--
		}
		if B == 1 {
			res += "b"
		}
		if bTimes > 0 {
			res += strings.Repeat("b", 2)
			bTimes--
		}
		if C == 1 {
			res += "c"
		}
		if cTimes > 0 {
			res += strings.Repeat("c", 2)
			cTimes--
		}
	}
	return res
}
