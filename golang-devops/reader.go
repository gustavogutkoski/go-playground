package main

import (
	"fmt"
	"io"
	"log"
)

type MyReader struct {
	contents string
	pos      int
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	if r.pos+1 <= len(r.contents) {
		n = copy(p, r.contents[r.pos:r.pos+1])
		r.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	myReader := MyReader{
		contents: "testing",
	}

	out, err := io.ReadAll(&myReader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)

}
