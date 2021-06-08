package main

import "fmt"

func main() {
	fmt.Println(add(100, 200))
	/*
		the above should print the follwing
		------
		Begin operation with x = 100 and y = 200
		Operation result = 300
		End peration
		300
	*/
}

/* The below function code is not accessible to you */
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
