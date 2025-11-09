package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	p2 := parse(i2)
	p3 := parse(i3)
	return dig(p1, false), dig(p2, false), dig(p3, true)
}

func dig(p *parsed, diagonal bool) int {
	offsets := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	if !diagonal {
		offsets = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	}
	res := p.s
	var done bool
	n := byte(2)
	for !done {
		done = true
		for y := 1; y < p.h-1; y++ {
		SQ:
			for x := 1; x < p.w-2; x++ {
				if p.m[y*p.w+x] != n-1 {
					continue
				}
				for _, offset := range offsets {
					if p.m[(y+offset[1])*p.w+(x+offset[0])] < n-1 {
						continue SQ
					}
				}
				p.m[y*p.w+x] = n
				res++
				done = false
			}
		}
		n++
	}
	return res
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	m    []byte
	w, h int
	s    int
}

func parse(in []byte) *parsed {
	w, h := 0, 0
	j := 0
	s := 0
	for i := range len(in) {
		switch in[i] {
		case '\n':
			h++
			w = j
			j = 0
		case '.':
			in[i] = 0
		case '#':
			s++
			in[i] = 1
		}
		j++
	}
	return &parsed{in, w, h + 1, s}
}

func (p *parsed) String() string {
	var sb strings.Builder
	for y := range p.h {
		for x := range p.w - 1 {
			sb.WriteByte(p.m[y*p.w+x] + '0')
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
