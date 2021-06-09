package main

import "fmt"

func main() {
	operations := map[int]func(int, int) int{
		1: add,
		2: subtract,
		3: multiply,
		4: divide,
	}
	for {
		choice := getMenuChoice()
		if operation, exists := operations[choice]; exists {
			no1, no2 := getOperands()
			result := operation(no1, no2)
			fmt.Println("Result = ", result)
		} else {
			return
		}
	}
}

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

func getMenuChoice() int {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	var input int
	fmt.Scanf("%d\n", &input)
	return input
}

func getOperands() (int, int) {
	fmt.Println("Enter the operands")
	var no1, no2 int
	fmt.Scanf("%d %d\n", &no1, &no2)
	return no1, no2
}
