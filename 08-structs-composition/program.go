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

type PerishableProduct struct { //Association
	product Product
	Expiry  string
}

type CompPerishableProduct struct {
	Product // composition because no variable name for Product given
	Expiry  string
}

func main() {

	var grapes = PerishableProduct{
		product: Product{102, "grape1", 10.0, 100, "fruits"},
		Expiry:  "2 days",
	}
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.product.Name)
	fmt.Println(ToString(&grapes.product))
	ApplyDiscount(&grapes.product, 10)
	fmt.Println("Prod pointer passed - so prod after making the discount permanent")
	fmt.Printf("%#v\n", grapes)

	var cGrapes = CompPerishableProduct{
		Product{102, "grape1", 10.0, 100, "fruits"},
		"2 days",
	}
	fmt.Printf("%#v\n", cGrapes)

	var cGrapes2 = CompPerishableProduct{
		Product: Product{102, "grape2", 10.0, 100, "fruits"},
		Expiry:  "2 days",
	}
	fmt.Printf("%#v\n", cGrapes2)
	fmt.Println("cGrapes2.Product.Name:", cGrapes2.Product.Name)
	fmt.Println("cGrapes2.Name:", cGrapes2.Name) //direct access

	ApplyDiscount(&cGrapes2.Product, 10)
	fmt.Println("After ApplyDiscount(&cGrapes2.Product, 10)")
	fmt.Printf("%#v\n", cGrapes2)

}

func ToString(product *Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q\n",
		product.Id, product.Name, product.Cost, product.Units, product.Category)
}
func ApplyDiscount(prod *Product, discount float64) float64 {
	prod.Cost *= (100 - discount) / 100
	return prod.Cost
}
