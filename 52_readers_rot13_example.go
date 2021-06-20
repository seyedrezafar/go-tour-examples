package main

import (
	"io"
	"os"
	"strings"
//	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	var len int

	if i, e := r.r.Read(b); e != io.EOF {
		len = i
		for j := 0; j < i; j++ {
			if b[j] < 'N' || (b[j] >= 'a' && b[j] < 'n') {
				b[j] += 13
			} else {
				b[j] -= 13
			}
		}
	} else {
		len = i
	}

	return len, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

