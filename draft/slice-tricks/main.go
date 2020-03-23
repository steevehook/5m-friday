package main

import "fmt"

// For more examples check out
// https://github.com/golang/go/wiki/SliceTricks
func main() {
	names := []string{"John", "Steve"}

	//push
	names = append(names, "Mike")

	// push front / unshift
	names = append([]string{"Mike"}, names...)

	// pop
	last, names := names[len(names)-1], names[:len(names)-1]
	fmt.Println("popped:", last)

	// pop front / shift
	first, names := names[0], names[1:]
	fmt.Println("shifted:", first)

	// append slice
	names = append(names, []string{"Chris", "Jane"}...)

	// copy
	namesCopy := make([]string, len(names))
	copy(namesCopy, names)
	// or
	namesCopy = append([]string(nil), names...)
	// or
	namesCopy = append(names[:0:0], names...)

	// cut
	names = append(names[:1], names[2:]...)
	fmt.Println("cut", names)

	// delete
	i := 1
	names = append(names[:i], names[i+1:]...)
	fmt.Println("delete", names)
	// or
	// names = names[:i+copy(names[i:], names[i+1:])]

	// insert slice at position
	pos := 1
	names = append(names[:pos], append([]string{"Amie", "Grace"}, names[pos:]...)...)
	fmt.Println("inserted at: ", pos, names)

	// extend
	names = append(names, make([]string, 10)...)
	fmt.Println("extended", len(names))
}
