package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
type PerishableProduct struct {
	product Product
	Expiry  string
}
*/
type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//var p Product
	//p := Product{}
	p := Product{Id: 100, Name: "Pen", Cost: 5, Units: 100, Category: "Stationary"}
	//p := Product{100, "Pen", 5, 100, "Stationary"}

	print(p)
	//applyDiscount(&p, 10)
	p.applyDiscount(10)
	print(p)

	//pp := PerishableProduct{product: Product{500, "Grapes", 50, 60, "Food"}, Expiry: "2 Days"}
	//pp := PerishableProduct{Product{500, "Grapes", 50, 60, "Food"}, "2 Days"}
	pp := NewPerishableProduct(500, "Grapes", 50, 60, "Food", "2 Days")
	fmt.Println(pp)
	fmt.Println("Cost of grapes => ", pp.Product.Cost)
}

func print(p Product) {
	fmt.Printf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
func applyDiscount(p *Product, discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}
*/

//applyDiscount as a method of Product
func (p *Product) applyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
