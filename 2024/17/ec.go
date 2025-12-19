package main

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	p2 := parse(i2)
	p3 := parse(i3)
	a3 := p3.part3()
	return p1.prims(math.MaxInt), p2.prims(math.MaxInt), a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type (
	star struct {
		x, y int
	}
	parsed struct {
		s []star
	}
)

func (p *parsed) part3() int {
	s := []int{}
	for len(p.s) > 0 {
		cs := p.prims(6)
		s = append(s, cs)
	}
	slices.SortFunc(s, func(a, b int) int {
		return b - a
	})
	return s[0] * s[1] * s[2]
}

func (p *parsed) prims(maxDist int) int {
	d := 0
	joined := []star{p.s[0]}
	unjoined := p.s[1:]
	for len(unjoined) > 0 {
		m := math.MaxInt
		mi := 0
		for i := 0; i < len(unjoined); i++ {
			for j := 0; j < len(joined); j++ {
				dm := Abs(unjoined[i].x-joined[j].x) + Abs(unjoined[i].y-joined[j].y)
				if dm >= maxDist {
					continue
				}
				if dm < m {
					m = dm
					mi = i
				}
			}
		}
		if m == math.MaxInt {
			break
		}
		joined = append(joined, unjoined[mi])
		unjoined[mi] = unjoined[len(unjoined)-1]
		unjoined = unjoined[:len(unjoined)-1]
		d += m
	}
	p.s = unjoined
	return d + len(joined)
}

func parse(in []byte) *parsed {
	p := &parsed{}
	w := bytes.IndexByte(in, '\n')
	h := (len(in) + 1) / (w + 1)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if in[x+y*(w+1)] == '*' {
				p.s = append(p.s, star{1 + x, h - y})
			}
		}
	}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
