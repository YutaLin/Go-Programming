package main

import "fmt"

func main() {
	mapInitByMake()
	mapInitByShorthand()
	// mapInitByVar()
	mapWithIfStatement()
	mapWithRange()
}

func mapInitByMake() {
	var myGreeting = make(map[string]string)
	myGreeting["Steven"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour."

	delete(myGreeting, "Jenny")

	fmt.Println(myGreeting)
}

func mapInitByShorthand() {

	myGreeting := map[string]string{}
	myGreeting["Steven"] = "Good Morning"
	myGreeting["Peter"] = "Aloha"

	myGreeting = map[string]string{
		"Steven": "Good Afternoon",
		"Peter":  "OHYA",
	}

	fmt.Println(myGreeting)
}

func mapInitByVar() {

	var myGreeting map[string]string

	myGreeting["Steven"] = "OH NO"
}

func mapWithIfStatement() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	delete(myGreeting, 2)

	if value, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("value:", value)
		fmt.Println("exists", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("value", value)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)

}

func mapWithRange() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for key, value := range myGreeting {
		fmt.Println(key, " - ", value)
	}
}
