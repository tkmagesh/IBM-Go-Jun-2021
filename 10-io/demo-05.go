/* using the bufio.NewReader to read sentences from the source */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("data2.dat")
	if fileError != nil {
		log.Fatalln(fileError)
	}
	defer fileHandle.Close()

	inputReader := bufio.NewReader(fileHandle)
	for {
		sentence, err := inputReader.ReadString('.')
		fmt.Println(sentence)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
	}
}
