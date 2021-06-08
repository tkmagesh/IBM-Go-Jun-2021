package main

import "fmt"

func main() {
	fmt.Println(add(100, 200))
	fmt.Println(subtract(100, 200))
	//fmt.Println(divide(40, 6))
	/*
		quotient, remainder := divide(40, 6)
		fmt.Printf("Dividing 40 by 6, quotient = %d, remainder = %d\n", quotient, remainder)
	*/
	//printing only the quotient ignoring the remainder
	quotient, _ := divide(40, 6)
	fmt.Printf("Dividing 40 by 6, quotient = %d\n", quotient)

}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) (result int) {
	result = x - y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
