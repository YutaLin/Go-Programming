package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)

	// Type
	fmt.Printf("%T\n", y)

	// Binary
	fmt.Printf("%b\n", y)

	// Hexidecimal
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// String print
	output := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(output)

	// The value of default format
	fmt.Printf("%v", y)
}
