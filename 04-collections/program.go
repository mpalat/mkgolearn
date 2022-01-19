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

	// Slices - Internally A pointer to an underlying array - but slice is dynamic
	fmt.Println("Slices")
	var products []string = []string{"pen", "pencil", "marker"}
	fmt.Println(products)

	var prods []string
	prods = append(prods, "Pen")
	fmt.Printf("prods cap:%d len:%d\n", cap(prods), len(prods))
	prods = append(prods, "Pencil", "Marker")
	fmt.Printf("prods cap:%d len:%d\n", cap(prods), len(prods))
	prods = append(prods, "Mouse")
	fmt.Printf("prods cap:%d len:%d\n", cap(prods), len(prods))
	fmt.Println(prods)

	var prodss []string = make([]string, 0, 10)
	prodss = append(prodss, "Pen")
	fmt.Printf("prodss cap:%d len:%d\n", cap(prodss), len(prodss))
	prodss = append(prodss, "Pencil", "Marker")
	fmt.Printf("prodss cap:%d len:%d\n", cap(prodss), len(prodss))
	prodss = append(prodss, "Mouse")
	fmt.Printf("prodss cap:%d len:%d\n", cap(prodss), len(prodss))
	fmt.Println(prodss)

	// copying
	nProds := prods
	fmt.Println(nProds)

	var prodStartEmpty []string = make([]string, 3, 10)
	fmt.Println("prodStartEmpty")
	fmt.Println(prodStartEmpty)
	prodStartEmpty = append(prodStartEmpty, "Pen")
	fmt.Printf("prodStartEmpty cap:%d len:%d\n", cap(prodStartEmpty), len(prodStartEmpty))
	prodStartEmpty = append(prodStartEmpty, "Pencil", "Marker")
	fmt.Printf("prodStartEmpty cap:%d len:%d\n", cap(prodStartEmpty), len(prodStartEmpty))
	prodStartEmpty = append(prodStartEmpty, "Mouse")
	fmt.Printf("prodStartEmpty cap:%d len:%d\n", cap(prodStartEmpty), len(prodStartEmpty))
	fmt.Println(prodStartEmpty)
	for idx, val := range prodStartEmpty {
		fmt.Printf("prodStartEmpty[%d] =%s\n", idx, val)
	}

}
