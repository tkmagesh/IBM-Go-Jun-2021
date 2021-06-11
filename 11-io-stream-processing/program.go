package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenResultCh := make(chan int)
	oddResultCh := make(chan int)
	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}

	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenResultCh, processWg)
	go sum(oddCh, oddResultCh, processWg)
	go merge(evenResultCh, oddResultCh, "result.dat", processWg)

	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Job Done!")

}

func source(filename string, ch chan int, wg *sync.WaitGroup) {
	fileHandle, fileError := os.Open(filename)
	if fileError != nil {
		panic(fileError)
	}
	defer func() {
		fileHandle.Close()
		wg.Done()
	}()
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		str := scanner.Text()
		val, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		ch <- val
	}
}

func splitter(in chan int, evenCh chan int, oddChan chan int, wg *sync.WaitGroup) {
	for no := range in {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddChan <- no
		}
	}
	close(evenCh)
	close(oddChan)
	wg.Done()
}

func sum(noCh chan int, resultCh chan int, wg *sync.WaitGroup) {
	result := 0
	for no := range noCh {
		result += no
	}
	resultCh <- result
	wg.Done()
}

func merge(evenResultCh chan int, oddResultCh chan int, resultFile string, wg *sync.WaitGroup) {
	resultFileHandle, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		resultFileHandle.Close()
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case evenResult := <-evenResultCh:
			resultFileHandle.Write([]byte(fmt.Sprintf("Even Total : %d\n", evenResult)))
		case oddResult := <-oddResultCh:
			resultFileHandle.Write([]byte(fmt.Sprintf("Odd Total : %d\n", oddResult)))
		}
	}
}
