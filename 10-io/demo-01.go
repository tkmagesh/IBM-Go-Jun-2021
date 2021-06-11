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

func Copy(writer io.Writer, reader io.Reader) (int64, error) {
	var totalBytes int64
	for {
		buffer := make([]byte, 1024)
		bytesRead, err := reader.Read(buffer)
		totalBytes += int64(bytesRead)
		if bytesRead > 0 {
			_, writerErr := writer.Write(buffer)
			if writerErr != nil {
				return totalBytes, writerErr
			}
		}
		if err != nil {
			return totalBytes, err
		}
	}
}

func main() {
	alphaReader := AlphaReader("Hi, How are you?")
	io.Copy(os.Stdout, alphaReader)
	/* Copy(os.Stdout, alphaReader) */
	fmt.Println("Done")
}
