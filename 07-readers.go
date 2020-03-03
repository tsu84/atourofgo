package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	fmt.Printf("%q",'A')
	return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}