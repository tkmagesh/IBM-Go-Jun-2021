package main

import "fmt"

func main() {

	var fn func() = func() {
		fmt.Println("fn is invoked")
	}
	fn()
	var add func(int, int) int = func(x, y int) int {
		return x + y
	}
	fmt.Printf("Type of add is %T\n", add)
	var result int = add(100, 200)
	fmt.Println(result)

	execute(fn)

}

func execute(fn func()) {
	fmt.Println("Start : Invoking the given function")
	fn()
	fmt.Println("End : Invoking the given function")
}
