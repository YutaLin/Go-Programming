package main

import (
	"fmt"
)

func main() {

	n := "Bond"

	switch n {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("miss money or bond or dr no")
		fallthrough
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
