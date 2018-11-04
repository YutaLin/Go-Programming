package main

import "fmt"

// Slices
func main() {
	// composite literal
	// A Slice allows to group together values of the same type
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// For loop
	for index, value := range x {
		println(index, value)
	}

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	printSlice(d)
	printSlice(e)
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	printSlice(d)
	printSlice(e)
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	// Append element to Slice
	mySlice := make([]int, 0, 5)
	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len:%d Cap:%d Value:%d\n", len(mySlice), cap(mySlice), mySlice[i])
	}

	// Initailization way 2
	initailizaition_way_2()
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v", len(s), cap(s), s)
}

// Slice Initailizaition Way 2

func initailizaition_way_2() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
