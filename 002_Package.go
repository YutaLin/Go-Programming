package main

import "fmt"

func main() {
	n, error := fmt.Println("one", 1, true)
	fmt.Println(n)
	fmt.Println(error)
}
