package main

import (
	"fmt"
)

// Go is a static progamming language

var y = 42
var z = "Shaken, not stirred"
var x string = `He said, "Shaken, not stirred"

`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
