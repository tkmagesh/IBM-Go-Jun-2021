//using the io.Reader

package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(buffer []byte) (int, error) {
	fmt.Println("Buffer size => ", len(buffer))
	count := 0
	for i := 0; i < len(alphaReader); i++ {
		if (alphaReader[i] >= 'A' && alphaReader[i] <= 'Z') || (alphaReader[i] >= 'a' && alphaReader[i] <= 'z') {
			buffer[count] = alphaReader[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	alphaReader := AlphaReader("Hi, How are you?")
	io.Copy(os.Stdout, alphaReader)
	fmt.Println("Done")
}
