//buffered channels

package main

import (
	"fmt"
)

func fn(ch chan int) {
	fmt.Println("[fn] before writing 10 into the channel")
	ch <- 10
	fmt.Println("[fn] after writing 10 into the channel")
	fmt.Println("[fn] before writing 20 into the channel")
	ch <- 20
	fmt.Println("[fn] after writing 20 into the channel")
	fmt.Println("[fn] before writing 30 into the channel")
	ch <- 30
	fmt.Println("[fn] after writing 30 into the channel")
	fmt.Println("[fn] before writing 40 into the channel")
	ch <- 40
	fmt.Println("[fn] after writing 40 into the channel")
	fmt.Println("[fn] before writing 50 into the channel")
	ch <- 50
	fmt.Println("[fn] after writing 50 into the channel")
}

func main() {
	ch := make(chan int, 2)
	go fn(ch)
	//time.Sleep(1 * time.Second)
	fmt.Println("[main] before reading data from the channel")
	fmt.Println("[main] data from channel ", <-ch)
	fmt.Println("[main] after reading data from the channel")
}
