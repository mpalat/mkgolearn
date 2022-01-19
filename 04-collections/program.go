package main

import "fmt"

func main() {

	// Arrays
	fmt.Println("Arrays")
	nos := [5]int{4, 1, 2, 5, 3}
	fmt.Println(nos)

	fmt.Println("Indexer iteration")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("range iteration")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// Copying
	newNos := nos
	for idx, val := range newNos {
		fmt.Printf("newNos[%d] = %d\n", idx, val)
	}
}
