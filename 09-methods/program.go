package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct { //Association
	Product
	Expiry string
}

func main() {
	product := Product{102, "grape1", 10.0, 100, "fruits"} // product is a value type
	product.ApplyDiscount(10)                              // here the product is a pointer
	fmt.Println("Prod pointer passed - so prod after making the discount permanent")
	fmt.Printf("%#v\n", product)
	product.jlt(10) // here the product is a value

	productPtr := &Product{102, "grape1", 10.0, 100, "fruits"}
	productPtr.ApplyDiscount(10)
	fmt.Println("Prod pointer passed - so prod after  the discount permanent")
	fmt.Printf("%#v\n", productPtr)

	var grapes = PerishableProduct{
		Product: Product{102, "grape1", 10.0, 100, "fruits"},
		Expiry:  "2 days",
	}
	fmt.Println("grapes.Product.ToString():\n", grapes.Product.ToString())
	fmt.Println(grapes.Product.ToString())
	grapes.ApplyDiscount(10)
	fmt.Println("After grapes.ApplyDiscount(10) - grapes.Product.ToString():\n", grapes.Product.ToString())
	fmt.Println("grapes.ToString():", grapes.ToString())

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

/* method overriding */
func (pp *PerishableProduct) ToString() string {
	return fmt.Sprintf("%s, Expiry=%s", pp.Product.ToString(), pp.Expiry)
}
