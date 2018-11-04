package main

import (
	"fmt"
)

func main() {

	// Example 1
	records := make([][]string, 0)

	// Student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	// Store the record
	records = append(records, student1)

	// Student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	// Store the record
	records = append(records, student2)

	// print
	fmt.Println(records)

	// Example 2

	transactions := make([][]int, 0)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)

		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
