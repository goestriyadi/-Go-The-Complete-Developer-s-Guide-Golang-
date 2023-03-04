package main

import (
	"fmt"
	"io"
	"os"
)

type logWritter struct {
}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {

	newFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWritter{}
	io.Copy(lw, newFile)
}
