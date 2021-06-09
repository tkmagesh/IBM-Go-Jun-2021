package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	//var p Product
	//p := Product{}
	p := Product{Id: 100, Name: "Pen", Cost: 5, Units: 100, Category: "Stationary"}
	print(p)
	applyDiscount(&p, 10)
	print(p)
}

func print(p Product) {
	fmt.Printf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func applyDiscount(p *Product, discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}
