package main

import (
	"fmt"
	"time"
)

//share memory for communication

func add(x, y int, ch chan int, delay time.Duration) {
	fmt.Println("Entering add with ", x, " and ", y)
	result := x + y
	//write data into the channel
	time.Sleep(delay * time.Second)
	fmt.Println("Writing result ", result)
	ch <- result
}

func main() {
	ch := make(chan int)
	go add(1000, 2000, ch, time.Duration(10))
	go add(100, 200, ch, time.Duration(5))
	//reading data from the channel
	fmt.Println("Waiting for result 1")
	result1 := <-ch
	fmt.Println("Waiting for result 2")
	result2 := <-ch
	fmt.Println(result1, result2)
}
