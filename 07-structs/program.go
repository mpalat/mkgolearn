package main

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	var product Product
	fmt.Println(product)
	fmt.Printf("%#v\n", product)

	var prod Product = Product{101, "Pen", 10, 100, "Stationary"}
	fmt.Printf("%#v\n", prod)

	// var prodPart Product = Product{
	// 	Id:       101,
	// 	Category: "Stationary",
	// }
	// fmt.Printf("%#v\n", prodPart)

	// prodP := Product{101, "Pen", 10, 100, "Stationary"}
	// fmt.Printf("%#v\n", ToString(prodP))

	// struct equality comparison by value and not by reference
	// prod == prodP

	var prodPtr *Product = &prod
	fmt.Printf("(*prodPtr).Id = %d\n", (*prodPtr).Id)
	fmt.Printf("prodPtr.Id = %d\n", prodPtr.Id)
	fmt.Println(prodPtr)

}

// func ToString(product Product) string {
// 	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q\n",
// 		product.Id, product.Name, product.Cost, product.Units, product.Category)
// }
func ToString(product *Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q\n",
		product.Id, product.Name, product.Cost, product.Units, product.Category)
}
