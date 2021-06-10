package main

import (
	"fmt"
	"time"
)

func print(msg string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		time.Sleep(delay)
		fmt.Println(msg)
	}
}

func main() {
	go print("Hello", 1*time.Second)
	go print("World", 3*time.Second)
	var input string
	fmt.Scanln(&input)
}
