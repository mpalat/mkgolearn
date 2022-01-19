package main

import "fmt"

func main() {

	// Slices - Internally A pointer to an underlying array - but slice is dynamic
	fmt.Println("Slices")
	var prods []string
	prods = append(prods, "Pen")
	prods = append(prods, "Pencil", "Marker")
	prods = append(prods, "Mouse")
	fmt.Println(prods)

	/*
	 [lo,hi)
	 [lo:)
	*/
	fmt.Println("2:5]", prods[2:5])
	fmt.Println("2:]", prods[2:5])
	fmt.Println(":5]", prods[2:5])

	newProds := prods[:] // another way to copy - but this is a  copy of the meta data
	fmt.Println("newProds", newProds)
	fmt.Println("adding Charger to newProds")
	newProds[0] = "Charger"
	fmt.Println("prods", prods)

	fmt.Println("Appending Adap to newProds")
	newProds = append(newProds, "Adapt")
	fmt.Println("newProds", newProds)
	fmt.Println("prods", prods)

	fmt.Println("Appending Adapter to prods")
	prods = append(prods, "Adapter")
	fmt.Println("newProds", newProds)
	fmt.Println("prods", prods)

	fmt.Println("Note only the header references are updated for prod and newProd - so be careful")
	fmt.Println("Maybe useful in in-array sorting where same parts of the same array to be manipulated")

	fmt.Println("Operation: prods = append(prods[0:2], prods[4:]...) ")
	prods = append(prods[0:2], prods[4:]...) // ... is the opertator to spread out the values
	fmt.Println(prods)
}
