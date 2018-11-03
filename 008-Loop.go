package main

import (
	"fmt"
)

// There's no while in loop

func main() {

	// for init; condition; post{ }
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// Nesting loops
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, j)
		}
	}

	// Single Condition
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	y := 1
	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y++
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}

	fmt.Println("done.")
}
