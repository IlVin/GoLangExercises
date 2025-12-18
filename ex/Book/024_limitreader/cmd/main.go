package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitreader struct {
	reader *io.Reader
	left   int
}

func LimitReader(r io.Reader, limit int) limitreader {
	return limitreader{reader: &r, left: limit}
}

func (lr *limitreader) Read(buf []byte) (int, error) {
	if lr.left <= 0 {
		return 0, io.EOF
	}

	buf = buf[0:lr.left]

	n, err := (*lr.reader).Read(buf)
	if err != nil {
		panic("")
	}
	lr.left -= n
	return n, err
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, &lr)
	if err != nil {
		log.Fatal(err)
	}
}
