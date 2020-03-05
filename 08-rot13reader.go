package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, e := rot13.r.Read(b)
	for i := 0; i < len(b); i++ {
		switch {
			case b[i] >= 65 && b[i] < 65+13:
				b[i] += 13
			case b[i] >= 97 && b[i] < 97+13:
				b[i] += 13
			case b[i] >= 65+13 && b[i] <= 65+13+13:
				b[i] -= 13
			case b[i] >= 97+13 && b[i] <= 97+13+13:
				b[i] -= 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
