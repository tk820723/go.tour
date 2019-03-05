package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	for i, ch := range b {
		if ch >= 65 && ch <= 77 || ch >= 97 && ch <= 109 {
			b[i] += 13
		} else if ch >= 78 && ch <= 90 || ch >= 110 && ch <= 122 {
			b[i] -= 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
