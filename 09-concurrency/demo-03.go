package main

import (
	"fmt"
	"sync"
)

//share memory for communication
var result = 0

var wg = &sync.WaitGroup{}
var mutex = &sync.Mutex{}

func add(x, y int) {
	mutex.Lock()
	result = x + y
	mutex.Unlock()
	wg.Done()
}

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println(result)
}
