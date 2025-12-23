package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, string) {
	p1 := parse(i1)
	p2 := parse(i2)
	p3 := parse(i3)
	mn, mx := p3.part3()
	return p1.part1(), p2.part2(), fmt.Sprint(mn, mx)
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	m [][]byte
	l [][]byte
}

func (p *parsed) part1() int {
	a := 0
	for i := range p.l {
		slot := p.drop(i, i+1)
		s := max(0, slot*2-(i+1))
		a += s
	}
	return a
}

func (p *parsed) part2() int {
	a := 0
	w := len(p.m[0])
	for i := range p.l {
		as := 0
		for s := 1; (s-1)*2 < w; s++ {
			slot := p.drop(i, s)
			as = max(as, slot*2-s)
		}
		a += as
	}
	return a
}

func (p *parsed) part3() (int, int) {
	w := len(p.m[0])
	scores := make([][]int, len(p.l))
	for i := range p.l {
		for s := 1; (s-1)*2 < w; s++ {
			slot := p.drop(i, s)
			sc := max(0, slot*2-s)
			scores[i] = append(scores[i], sc)
		}
	}
	Dg(scores)
	mx := math.MinInt
	mn := math.MaxInt
	c := make([]int, len(p.l))
OUTER:
	for {
		sc := 0
		valid := true
	LOOP:
		for j := 0; j < len(c); j++ {
			for k := j + 1; k < len(c); k++ {
				if c[j] == c[k] {
					valid = false
					break LOOP
				}
			}
			sc += scores[j][c[j]]
		}
		if valid {
			mx = max(mx, sc)
			mn = min(mn, sc)
		}
		for j := 0; j < len(c); j++ {
			c[j]++
			if c[j] < len(scores[0]) {
				break
			}
			c[j] = 0
			if j == len(c)-1 {
				break OUTER
			}
		}
	}
	return mn, mx
}

func (p *parsed) drop(li, slot int) int {
	x := (slot - 1) * 2
	l := p.l[li]
	w := len(p.m[0])
	i := 0
	for y := range p.m {
		switch p.m[y][x] {
		case '*':
			if l[i] == 'R' {
				x++
				if x >= w {
					x = w - 2
				}
			} else {
				x--
				if x < 0 {
					x = 1
				}
			}
			i++
		case '.':
			// noop
		}
	}
	return 1 + x/2
}

func parse(in []byte) *parsed {
	p := &parsed{}
	i := bytes.Index(in, []byte{'\n', '\n'})
	p.m = bytes.Split(in[:i], []byte{'\n'})

	i += 2
	p.l = bytes.Split(in[i:], []byte{'\n'})
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
