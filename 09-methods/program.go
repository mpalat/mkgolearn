package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	product := Product{102, "grape1", 10.0, 100, "fruits"} // product is a value type
	product.ApplyDiscount(10)                              // here the product is a pointer
	product.jlt(10)                                        // here the product is a value

	productPtr := &Product{102, "grape1", 10.0, 100, "fruits"}
	productPtr.ApplyDiscount(10)

}

func (product *Product) ToString() string { // make product as a receiver
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q\n",
		product.Id, product.Name, product.Cost, product.Units, product.Category)
}
func (prod *Product) ApplyDiscount(discount float64) float64 {
	prod.Cost *= (100 - discount) / 100
	return prod.Cost
}
func (prod Product) jlt(discount float64) float64 {
	return prod.Cost
}
