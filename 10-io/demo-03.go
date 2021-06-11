//implementing a writer

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type ChannelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *ChannelWriter {
	return &ChannelWriter{
		Channel: make(chan byte, 1024),
	}
}

func (cw *ChannelWriter) Write(buffer []byte) (int, error) {
	if len(buffer) == 0 {
		return 0, nil
	}
	go func() {
		defer close(cw.Channel)
		for _, b := range buffer {
			cw.Channel <- b
		}
	}()
	return len(buffer), nil
}

func main() {
	cw := NewChannelWriter()
	fileReader, err := os.Open("data1.dat")
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		io.Copy(cw, fileReader)
	}()
	for c := range cw.Channel {
		fmt.Printf("%c\n", c)
	}

}
