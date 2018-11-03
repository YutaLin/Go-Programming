package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	// Initailization statement
	if x := 42; x == 42 {
		fmt.Println("001")
	}

	// If Else statement
	y := 42
	if y == 42 {
		fmt.Println("x is 42")
	} else {
		fmt.Println("x is not 42")
	}
}
