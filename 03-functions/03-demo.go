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

	adder := getAdder()
	fmt.Println(adder(1000, 2000))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))
}

func execute(f func()) {
	fmt.Println("Start : Invoking the given function")
	f()
	fmt.Println("End : Invoking the given function")
}

func getAdder() func(int, int) int {
	add := func(x int, y int) int {
		return x + y
	}
	return add
}

func getAdderFor(x int) func(int) int {
	add := func(y int) int {
		return x + y
	}
	return add
}
