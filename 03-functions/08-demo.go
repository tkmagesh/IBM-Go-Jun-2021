package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("A painc is raised and the application is exiting gracefully")
			fmt.Println(r)
			return
		} else {
			fmt.Println("normal exit")
		}
	}()
	result := divide(100, 0)
	fmt.Println("Result = ", result)
}

func divide(x, y int) int {
	defer func() {
		fmt.Println("deferred from divide")
	}()
	if y == 0 {
		panic("Invalid argument(s). Cannot divide by 0!")
	}
	return x / y
}
