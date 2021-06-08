//defer - demo
package main

import "fmt"

func main() {
	f1()
}

func f1() {

	defer func() {
		fmt.Println("[@defer-1] exiting f1")
	}()
	defer func() {
		fmt.Println("[@defer-2] exiting f1")
	}()
	defer func() {
		fmt.Println("[@defer-3] exiting f1")
	}()

	//defer fmt.Println("exiting f1")
	fmt.Println("processing f1")
	f2()
}

func f2() {

	defer func() {
		fmt.Println("[@defer-1] exiting f2")
	}()
	defer func() {
		fmt.Println("[@defer-2] exiting f2")
	}()
	defer func() {
		fmt.Println("[@defer-3] exiting f2")
	}()

	//defer fmt.Println("exiting f2")
	fmt.Println("processing f2")
}
