package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Hobby string `json:"hobby,omitempty"`
	Email string `json:"-"`
	Money float64 `json:"money,string"`
}

func main() {
	p1 := person{Name: "John", Hobby: "Music"}
	bs1, _ := json.Marshal(p1)
	fmt.Println(string(bs1))

	p2 := person{Name: "Jane", Email: "e@d.com"}
	bs2, _ := json.Marshal(p2)
	fmt.Println(string(bs2))

	p3 := person{Name: "Steve", Money: 200}
	bs3, _ := json.Marshal(p3)
	fmt.Println(string(bs3))
}
