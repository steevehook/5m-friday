package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Steve",
		Age:  25,
	}
	bs, _ := json.Marshal(res)
	fmt.Println(string(bs))
}
