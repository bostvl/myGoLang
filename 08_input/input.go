package main

import (
	"fmt"
	"io"
	"os"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	// TODO: add solution here
	p := make([]byte, 128)
	s := ""
	for {
		n, err := r.Read(p)
		if err != nil {
			if err == io.EOF {
				s = s + string(p[:n])
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		s = s + string(p[:n])
	}
	return s
}
