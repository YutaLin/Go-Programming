package main

import "fmt"

// Slices
func main() {
	// composite literal
	// A Slice allows to group together values of the same type
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// For loop
	for index, value := range x {
		println(index, value)
	}

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	printSlice(d)
	printSlice(e)
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	printSlice(d)
	printSlice(e)
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
