package main

import (
	"bytes"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.count(0, 0)
	p2 := parse(i2)
	a2 := p2.count(0, 0)
	a2 += p2.count(p2.w-2, p2.h-1)
	p3 := parse(i3)
	a3 := p3.part3()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	in   []byte
	w, h int
}

func parse(in []byte) *parsed {
	w := 1 + bytes.IndexByte(in, '\n')
	h := (1 + len(in)) / w
	p := &parsed{in, w, h}
	return p
}

func (p *parsed) get(x, y int) byte {
	if 0 <= x && x < p.w-1 && 0 <= y && y < p.h {
		return p.in[x+y*p.w]
	}
	return '@'
}

type point struct{ x, y int }

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type area struct {
	ch   byte
	p    []*point
	n    []*area
	used int
	done bool
}

func (a *area) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%p: %c: %v\n", a, a.ch, a.p)
	if len(a.n) > 0 {
		fmt.Fprint(&sb, "  ")
		for _, n := range a.n {
			fmt.Fprintf(&sb, "%p ", n)
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func (p *parsed) part3() int {
	g := make([]*area, (p.w-1)*p.h)
	var add func(x, y int, a *area)
	add = func(x, y int, a *area) {
		gi := x + y*(p.w-1)
		if g[gi] != nil {
			return
		}
		g[x+y*(p.w-1)] = a
		a.p = append(a.p, &point{x, y})
		if a.ch == p.get(x, y-1) {
			add(x, y-1, a)
		}
		if a.ch == p.get(x+1, y) {
			add(x+1, y, a)
		}
		if a.ch == p.get(x, y+1) {
			add(x, y+1, a)
		}
		if a.ch == p.get(x-1, y) {
			add(x-1, y, a)
		}
	}
	for y := range p.h {
		for x := range p.w - 1 {
			gi := x + y*(p.w-1)
			if g[gi] != nil {
				continue
			}
			ch := p.get(x, y)
			a := &area{ch: ch}
			add(x, y, a)
		}
	}
	var areas []*area
	for _, a := range g {
		if a.n != nil {
			continue
		}
		areas = append(areas, a)
		n := make(map[*area]*area)
		for _, xy := range a.p {
			x, y := xy.x, xy.y
			ch := p.get(x, y)
			if p.get(x, y+1) < ch {
				na := g[x+(y+1)*(p.w-1)]
				n[na] = na
			}
			if p.get(x+1, y) < ch {
				na := g[(x+1)+y*(p.w-1)]
				n[na] = na
			}
			if p.get(x, y-1) < ch {
				na := g[x+(y-1)*(p.w-1)]
				n[na] = na
			}
			if p.get(x-1, y) < ch {
				na := g[(x-1)+y*(p.w-1)]
				n[na] = na
			}
		}
		a.n = slices.Collect(maps.Values(n))
	}
	var count func(a *area, used int) int
	count = func(a *area, used int) int {
		if a.done {
			return 0
		}
		if a.used == used {
			return 0
		}
		a.used = used
		c := len(a.p)
		for _, n := range a.n {
			c += count(n, used)
		}
		return c
	}
	var mark func(a *area)
	mark = func(a *area) {
		if a.done {
			return
		}
		a.done = true
		for _, n := range a.n {
			mark(n)
		}
	}
	used := 1
	mx := math.MinInt
	mi := 0
	for i, a := range areas {
		c := count(a, used)
		if c > mx {
			mx = c
			mi = i
		}
		used++
	}
	res := mx
	mark(areas[mi])
	mx = math.MinInt
	mi = 0
	for i, a := range areas {
		c := count(a, used)
		if c > mx {
			mx = c
			mi = i
		}
		used++
	}
	res += mx
	mark(areas[mi])
	mx = math.MinInt
	mi = 0
	for i, a := range areas {
		c := count(a, used)
		if c > mx {
			mx = c
			mi = i
		}
		used++
	}
	res += mx
	mark(areas[mi])
	return res
}

func (p *parsed) count(x, y int) int {
	ch := p.get(x, y)
	if ch == '@' {
		return 0
	}
	p.in[x+y*p.w] = '@'
	c := 1
	if p.get(x-1, y) <= ch {
		c += p.count(x-1, y)
	}
	if p.get(x+1, y) <= ch {
		c += p.count(x+1, y)
	}
	if p.get(x, y-1) <= ch {
		c += p.count(x, y-1)
	}
	if p.get(x, y+1) <= ch {
		c += p.count(x, y+1)
	}
	return c
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, string(p.in))
	return sb.String()
}
