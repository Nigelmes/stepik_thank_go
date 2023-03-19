package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

type randomReader struct {
	max int
}

func (r *randomReader) Read(p []byte) (n int, err error) {
	if r.max <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.max {
		p = p[:r.max]
	}
	for i := range p {
		p[i] = byte(rand.Intn(256))
	}
	r.max -= len(p)
	return len(p), nil
}

func RandomReader(max int) io.Reader {
	return &randomReader{max: max}
}

// конец решения

func main() {
	rand.Seed(0)

	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
}
