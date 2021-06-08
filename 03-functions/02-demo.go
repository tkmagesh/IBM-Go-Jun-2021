//variadic function
package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
}

func sum(nos ...int) int {
	fmt.Println(len(nos))
	result := 0
	for index := 0; index < len(nos); index++ {
		result += nos[index]
	}
	return result
}
