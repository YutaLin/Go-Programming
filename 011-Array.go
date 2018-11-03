package main

import (
	"fmt"
)

// Golang doesn't recommand use array
func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
