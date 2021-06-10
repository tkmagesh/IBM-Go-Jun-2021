//generate fixed number of fibonacci series

package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, count int) {
	x, y := 0, 1
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		ch <- x
		x, y = y, x+y
		count -= 1
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(ch, 10)
	for no := range ch {
		fmt.Println(no)
	}
}
