//chaining io

package main

import (
	"io"
	"log"
	"os"
)

type AlphaReader struct {
	Src io.Reader
}

func (alphaReader AlphaReader) Read(buffer []byte) (int, error) {
	if len(buffer) == 0 {
		return 0, nil
	}
	inputData := make([]byte, len(buffer))
	count, err := alphaReader.Src.Read(inputData)
	if err != nil {
		return count, err
	}

	dataCount := 0
	for i := 0; i < len(inputData); i++ {
		if (inputData[i] >= 'A' && inputData[i] <= 'Z') || (inputData[i] >= 'a' && inputData[i] <= 'z') {
			buffer[dataCount] = inputData[i]
			dataCount++
		}
	}
	return dataCount, io.EOF
}

func main() {
	//strReader := strings.NewReader("Hi, How are you?")
	fileReader, err := os.Open("data1.dat")
	if err != nil {
		log.Fatalln(err)
	}
	//alphaReader := AlphaReader{Src: strReader}
	alphaReader := AlphaReader{Src: fileReader}
	io.Copy(os.Stdout, alphaReader)
}
