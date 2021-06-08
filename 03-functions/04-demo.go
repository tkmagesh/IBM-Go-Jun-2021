package main

import "fmt"

func main() {
	/*
		loggerAdd := func(x, y int) int {
			fmt.Printf("Begin operation with x = %d and y = %d\n", x, y)
			result := add(x, y)
			fmt.Printf("Operation result = %d\n", result)
			fmt.Println("End peration")
			return result
		}
		fmt.Println(loggerAdd(100, 200))

		loggerSubtract := func(x, y int) int {
			fmt.Printf("Begin operation with x = %d and y = %d\n", x, y)
			result := subtract(x, y)
			fmt.Printf("Operation result = %d\n", result)
			fmt.Println("End peration")
			return result
		}
		fmt.Println(loggerSubtract(100, 200))
	*/
	loggerAdd := loggerOperation(add)
	loggerSubtract := loggerOperation(subtract)

	fmt.Println(loggerAdd(100, 200))
	fmt.Println(loggerSubtract(100, 200))
	/*
		the above should print the follwing
		------
		Begin operation with x = 100 and y = 200
		Operation result = 300
		End peration
		300
	*/
}

func loggerOperation(operation func(x, y int) int) func(x, y int) int {
	return func(x, y int) int {
		fmt.Printf("Begin operation with x = %d and y = %d\n", x, y)
		result := operation(x, y)
		fmt.Printf("Operation result = %d\n", result)
		fmt.Println("End peration")
		return result
	}
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
