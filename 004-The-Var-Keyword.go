package main

import (
	"fmt"
)

var y = 34

// Declare there is a variable with identifier "z"
// "z" is a type of Int and it's value is 0
var z int

func main() {
	// short declaration operator
	// Declare a variable and assign a value
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again", y)
}
