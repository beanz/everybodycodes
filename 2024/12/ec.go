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
	var a1 int
	{
		for _, t := range p1.targets {
			a1 += p1.power(t)
		}
	}

	p2 := parse(i2)
	var a2 int
	{
		for _, t := range p2.targets {
			a2 += p2.power(t)
		}
		for _, t := range p2.hardTargets {
			a2 += 2 * p2.power(t)
		}
	}
	p3 := parse3(i3)
	var a3 int
	for _, t := range p3.targets {
		a := p3.power3(t)
		a3 += a
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
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

func parse3(in []byte) *parsed {
	p := &parsed{segment: []*point{{1, 1}, {1, 2}, {1, 3}}}
	for i := 0; i < len(in); i++ {
		j, x := ChompUInt[int](in, i)
		j, y := ChompUInt[int](in, j+1)
		i = j
		t := point{x, y}
		p.targets = append(p.targets, &t)
	}
	return p
}

func parse(in []byte) *parsed {
	p := &parsed{segment: make([]*point, 3)}
	w := bytes.IndexByte(in, '\n') + 1
	h := len(in) / w
	for y := h - 1; y >= 0; y-- {
		for x := 0; x < w-1; x++ {
			i := x + (h-y)*w
			t := point{x, y}
			switch in[i] {
			case 'A', 'B', 'C':
				p.segment[in[i]-'A'] = &t
			case 'T':
				p.targets = append(p.targets, &t)
			case 'H':
				p.hardTargets = append(p.hardTargets, &t)
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

func (p *parsed) power3(m *point) int {
	x := m.x / 2
	y := m.y - x - m.x&0x1
	for i, s := range p.segment {
		dx := x - s.x + 1
		dy := y - s.y + 1
		if dx < dy {
			continue
		}

		if dx <= 2*dy {
			return dy * (i + 1)
		}
		if px3 := dx + dy; px3%3 == 0 {
			return (px3 / 3) * (i + 1)
		}
	}
	return -9999
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "S: %v\nT: %v\nH: %v\n", p.segment, p.targets, p.hardTargets)
	return sb.String()
}
