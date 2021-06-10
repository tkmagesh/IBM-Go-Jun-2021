//generate fixed number of fibonacci series

package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, stop chan string) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}
}

func main() {
	fmt.Println("Press ENTER to stop...")
	ch := make(chan int)
	stop := make(chan string)
	go fibonacci(ch, stop)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	var input string
	fmt.Scanln(&input)
	stop <- "done"
	fmt.Println("Done")
}
