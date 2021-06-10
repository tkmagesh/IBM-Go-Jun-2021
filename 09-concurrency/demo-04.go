package main

import (
	"fmt"
)

//share memory for communication

func add(x, y int, ch chan int) {
	result := x + y
	//write data into the channel
	ch <- result
	ch <- 100
}

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	//reading data from the channel
	result := <-ch
	fmt.Println(result)
}
