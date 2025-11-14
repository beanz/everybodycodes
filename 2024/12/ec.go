package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	Lg(p1)
	var a1 int
	{
		for _, t := range p1.targets {
			a1 += p1.power(t)
		}
	}

	p2 := parse(i2)
	Lg(p2)
	var a2 int
	{
		for _, t := range p2.targets {
			a2 += p2.power(t)
		}
		for _, t := range p2.hardTargets {
			a2 += 2 * p2.power(t)
		}
	}
	return a1, a2, 0
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type point struct {
	x, y int
}

func (p *point) String() string {
	return fmt.Sprint(p.x, ",", p.y)
}

type parsed struct {
	targets     []*point
	hardTargets []*point
	segment     []*point
}

func parse(in []byte) *parsed {
	p := &parsed{segment: make([]*point, 3)}
	w := bytes.IndexByte(in, '\n') + 1
	h := len(in) / w
	for y := h - 1; y >= 0; y-- {
		for x := 0; x < w-1; x++ {
			i := x + (h-y)*w
			switch in[i] {
			case 'A', 'B', 'C':
				p.segment[in[i]-'A'] = &point{x, y}
			case 'T':
				p.targets = append(p.targets, &point{x, y})
			case 'H':
				p.hardTargets = append(p.hardTargets, &point{x, y})
			}
		}
	}
	return p
}

func (p *parsed) power(t *point) int {
	mn := math.MaxInt
	for i, s := range p.segment {
		dx := t.x - s.x
		dy := t.y - s.y
		if dx == dy {
			mn = min(mn, dx*(i+1))
		}
		if dy > 0 && dx <= dy && dy <= 2*dx {
			mn = min(mn, dy*(i+1))
		}
		if px3 := dx + dy; px3%3 == 0 && dy <= dx {
			mn = min(mn, (px3/3)*(i+1))
		}
	}
	return mn
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "S: %v\nT: %v\nH: %v\n", p.segment, p.targets, p.hardTargets)
	return sb.String()
}
