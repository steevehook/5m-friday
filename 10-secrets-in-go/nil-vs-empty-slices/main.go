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
	// slice zero value = nil
	p := people{
		{"John"},
		{"Steve"},
	}
	bs1, _ := json.Marshal(response{Items: p})
	bs2, _ := json.Marshal(response{}) // Items = nil
	bs3, _ := json.Marshal(response{Items: people(nil)}) // Items = nil
	bs4, _ := json.Marshal(response{Items: people{}}) // Items = empty
	fmt.Printf("%s\n%s\n%s\n%s", string(bs1), string(bs2), string(bs3), string(bs4))
}
