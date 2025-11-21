package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for rr := range p[:n] {
		if p[rr] >= 65 && p[rr] <= 77 || p[rr] >= 97 && p[rr] <= 109 {
			p[rr] = p[rr] + 13
		} else if p[rr] >= 78 && p[rr] <= 90 || p[rr] >= 110 && p[rr] <= 122 {
			p[rr] = p[rr] - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
