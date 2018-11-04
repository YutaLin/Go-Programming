package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42, 9}
	fmt.Println(x[2])
	fmt.Println(x[3:])
	fmt.Println(x[:4])
	fmt.Println(x[3:5])
}
