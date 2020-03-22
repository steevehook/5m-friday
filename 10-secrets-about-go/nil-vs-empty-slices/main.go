package main

import (
	"encoding/json"
	"fmt"
)

type people []struct {
	Name string `json:"name"`
}
type response struct {
	Items people `json:"items"`
}

func main() {
	p := people{
		{"John"},
		{"Steve"},
	}
	res := response{Items: p}
	bs1, _ := json.Marshal(res)
	bs2, _ := json.Marshal(response{}) // zero value
	bs3, _ := json.Marshal(response{Items: people{}})
	fmt.Printf("%s\n%s\n%s", string(bs1), string(bs2), string(bs3))
}
