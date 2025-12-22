package main

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	for range 100 {
		p1.iter()
	}
	p2 := parse(i2)
	p3 := parse(i3)
	a3 := p3.part2()
	return p1.score(), p2.part2(), a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type (
	point struct {
		x, y int
	}
	parsed struct {
		p []*point
		w int
	}
)

func parse(in []byte) *parsed {
	p := &parsed{}
	for i := 0; i < len(in); i++ {
		var x, y int
		i, x = ChompUInt[int](in, i+2)
		i, y = ChompUInt[int](in, i+3)
		p.p = append(p.p, &point{x, y})
		p.w = max(p.w, x+y)
	}
	return p
}

func (p *parsed) part2() int {
	a, d := 0, 1
	for _, e := range p.p {
		ep := e.x + e.y - 1
		for (e.x+a)%ep != 0 {
			a += d
		}
		d *= ep
	}
	return a
}

func (p *parsed) score() int {
	s := 0
	for _, e := range p.p {
		s += e.x + 100*e.y
	}
	return s
}

func (p *parsed) iter() bool {
	c := 0
	for _, e := range p.p {
		if e.y == 1 {
			e.x, e.y = e.y, e.x
		} else {
			e.y--
			e.x++
			if e.y == 1 {
				c++
			}
		}
	}
	return c == len(p.p)
}

func (p *parsed) String() string {
	var sb strings.Builder
	for y := 1; y < p.w; y++ {
		for x := 1; x < p.w-y+1; x++ {
			if slices.ContainsFunc(p.p, func(e *point) bool {
				return e.x == x && e.y == y
			}) {
				sb.WriteByte('@')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
