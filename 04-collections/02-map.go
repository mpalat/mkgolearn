// Map
package main

import "fmt"

func main() {

	prodRanks := map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 5,
	}
	// var prodRanks map[string]int = make(map[string]int)
	// prodRanks["Pen"] = 4
	fmt.Println(prodRanks)
	prodRanks["Pen"] = 34
	fmt.Println(prodRanks)

	fmt.Println("Iterating throug the map")
	for key, val := range prodRanks {
		fmt.Printf("key:%s, val:%d\n", key, val)
	}

	fmt.Println("Checking if a key exists")
	names := [2]string{"Pen", "Table"}
	for _, prodToSearch := range names {
		if val, exists := prodRanks[prodToSearch]; exists {
			fmt.Printf("%s exists:%d\n", prodToSearch, val)
		} else {
			fmt.Printf("%s does not exist\n", prodToSearch)
		}
	}

	fmt.Println("Deleting Pen")
	delete(prodRanks, "Pen")
	fmt.Println(prodRanks)
}
