package main

import (
	"fmt"
	"time"
)

func print(msg string, delay time.Duration, in chan string, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(delay)
		fmt.Println(msg)
		out <- "done"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", 1*time.Second, ch1, ch2)
	go print("World", 3*time.Second, ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}
