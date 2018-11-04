package main

import (
	"fmt"
)

func main() {
	createSliceByShorthand()
	createSliceByVar()
	createSliceByMake()
}

// Has underlying array but space and length are zero
// Should use append to add element

func createSliceByShorthand() {
	println("----Shorthand----")
	student := []string{}
	students := [][]string{}

	// Index out of range
	// student[0] = "Yuta"
	// Don't have space yet, need use append
	student = append(student, "Yuta")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	println("----Shorthand----")
}

// Doesn't have underlying array means nil
// Should us append to add element

func createSliceByVar() {
	println("-----Var----")
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	println("-----Var----")
}

// Has underlying array with given length and space
// Could use index within length to fetch/update element
// Coudld use append to add element

func createSliceByMake() {
	println("-----Make----")

	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	println("-----Make----")
}
