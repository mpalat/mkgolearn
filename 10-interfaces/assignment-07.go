/* Assignment:7 implement Sorting */
/*
Implement "sorting" functionality for the Products
	The user should be able to sort the products by any attribute
	Important: use the "sort.Sort()" function
*/

package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (product *Product) ToString() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type ProductPredicate func(product Product) bool

type Products []Product

func (products *Products) IndexOf(product Product) int {
	for idx, p := range *products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) >= 0
}

func (products *Products) Filter(criteria ProductPredicate) Products {
	filteredProducts := Products{}
	for _, p := range *products {
		if criteria(p) {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts
}

func (products *Products) All(criteria ProductPredicate) bool {
	for _, p := range *products {
		if !criteria(p) {
			return false
		}
	}
	return true
}

func (products *Products) Any(criteria ProductPredicate) bool {
	for _, p := range *products {
		if criteria(p) {
			return true
		}
	}
	return false
}

func (products *Products) ToString() string {
	result := ""
	for _, p := range *products {
		result += p.ToString() + "\n"
	}
	return result
}

type CompareFunc func(this, other Product) int

func idCompare(this, other Product) int {
	if this.Id < other.Id {
		return -1
	} else if this.Id > other.Id {
		return 1
	}
	return 0
}

func unitsCompare(this, other Product) int {
	if this.Units < other.Units {
		return -1
	} else if this.Units > other.Units {
		return 1
	}
	return 0
}

var compareFunc CompareFunc = idCompare // default assignment to id compare

func (Products Products) Len() int {
	return len(Products)
}

func (Products Products) Less(i, j int) bool {
	return compareFunc(Products[i], Products[j]) < 0
}
func (Products Products) Swap(i, j int) {
	Products[i], Products[j] = Products[j], Products[i]
}

func (Products *Products) sort(comparator CompareFunc) *Products {
	compareFunc = comparator
	sort.Sort(Products)
	return Products
}

func main() {
	var stove = Product{102, "Stove", 5000, 5, "Utencil"}
	var nonExistentProduct = Product{1000, "Dummy", 1000, 1000, "Dummy"}
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	products.sort(idCompare)
	fmt.Println(products)

	fmt.Println("Sorting on units")
	products.sort(unitsCompare)
	fmt.Println(products)

	fmt.Println("Index of stove = ", products.IndexOf(stove))
	fmt.Println("Index of nonExistentProduct = ", products.IndexOf(nonExistentProduct))

	fmt.Println("stove exists ? = ", products.Includes(stove))
	fmt.Println("nonExistentProduct exists ? = ", products.Includes(nonExistentProduct))

	fmt.Println("Filter")
	/*
		fmt.Println("Filter costly products")
		costlyProducts := products.FilterCostlyProdoucts()
		fmt.Print(costlyProducts.ToString())
		fmt.Println("Filter Stationary products")
		stationaryProducts := products.FilterStationaryProdoucts()
		fmt.Print(stationaryProducts.ToString())
	*/

	fmt.Println("Filter costly products")
	costlyProductCriteria := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductCriteria)
	fmt.Print(costlyProducts.ToString())

	fmt.Println("Filter Stationary products")
	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductCriteria)
	fmt.Print(stationaryProducts.ToString())

	fmt.Println("All")
	fmt.Println("Are all the products costly products ? :", products.All(costlyProductCriteria))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ? :", products.Any(costlyProductCriteria))
}
