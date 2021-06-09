package main

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	products := make([]Product, 5)
	products[0] = Product{101, "Pen", 5, 50, "Stationary"}
	products[1] = Product{105, "Pencil", 2, 100, "Stationary"}
	products[2] = Product{102, "Mouse", 100, 20, "Electronics"}
	products[3] = Product{104, "Marker", 30, 80, "Stationary"}
	products[4] = Product{103, "Charger", 200, 10, "Electronics"}

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
