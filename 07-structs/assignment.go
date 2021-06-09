package main

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//alias for []Product
type Products []Product

func main() {
	products := make(Products, 5)
	products[0] = Product{101, "Pen", 5, 50, "Stationary"}
	products[1] = Product{105, "Pencil", 2, 100, "Stationary"}
	products[2] = Product{102, "Mouse", 100, 20, "Electronics"}
	products[3] = Product{104, "Marker", 30, 80, "Stationary"}
	products[4] = Product{103, "Charger", 200, 10, "Electronics"}

	stylus := Product{107, "Stylus", 200, 10, "Stationary"}
	products.AddProduct(stylus)

	fmt.Println("Initial List")
	fmt.Println(products.Format())

	fmt.Println("Index of stylus => ", products.IndexOf(stylus))
	fmt.Println("Includes stylus ? => ", products.Includes(stylus))

	costlyProductPredicate := func(product Product) bool {
		return product.Cost >= 100
	}
	anyCostlyProducts := products.Any(costlyProductPredicate)
	fmt.Println("Any costly products => ", anyCostlyProducts)

	areAllCostlyProducts := products.All(costlyProductPredicate)
	fmt.Println("Are all costly products => ", areAllCostlyProducts)

	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println("Costly Products List")
	fmt.Println(costlyProducts.Format())
}

//write the functions for the following
/*
	1. AddProduct -> to add a new product to the products
	2. IndexOf -> returns the index of the given product in the products list
	3. Includes -> return true/false based on the presence of the given product in the products list
	4. Any -> return true/false based if there is atleast one product satisfying the given condition in the product list
	5. All -> return true/false based if there is all the products in the list satisfy the given condition
	6. Filter -> return the list of products that satisfy the given condition in the products list
*/

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (products Products) Format() string {
	result := ""
	for _, product := range products {
		result += product.Format()
	}
	return result
}

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}

func (products Products) IndexOf(product Product) int {
	for idx, item := range products {
		if item == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, item := range products {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, item := range products {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, item := range products {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
