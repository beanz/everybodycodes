package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.solve()
	p2 := parse(i2)
	a2 := p2.solve2()
	p3 := parse(i3)
	a3 := p3.solve2()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	m    map[int][][2]int
	w, h int
}

func parse(in []byte) *parsed {
	p := &parsed{m: make(map[int][][2]int)}
	for i := 0; i < len(in); i++ {
		var x, y, l int
		i, x = ChompUInt[int](in, i)
		i++
		i, y = ChompUInt[int](in, i)
		i++
		i, l = ChompUInt[int](in, i)
		p.m[x] = append(p.m[x], [2]int{y, y + l})
		p.w = max(p.w, x)
		p.h = max(p.h, y+l)
	}
	p.w++
	p.h += 8
	return p
}

func (p *parsed) solve2() int {
	var res int
	for x, wall := range p.m {
		f := (x + wall[0][0] + 1) / 2
		res = max(res, f)
	}
	return res
}

func (p *parsed) solve() int {
	type rec struct {
		x, y, f int
	}
	seen := map[rec]struct{}{}
	work := [1300][]rec{{{0, 0, 0}}}
	add := func(x, y, f int) {
		if y < 0 {
			return
		}
		r := rec{x, y, f}
		if _, ok := seen[r]; ok {
			return
		}
		seen[r] = struct{}{}
		work[f] = append(work[f], rec{x, y, f})
	}
	for qi := 0; qi < len(work); qi++ {
		for qj := 0; qj < len(work[qi]); qj++ {
			cur := work[qi][qj]
			if cur.x == p.w {
				return cur.f
			}
			walls := p.m[cur.x+1]
			if walls == nil {
				add(cur.x+1, cur.y-1, cur.f)
				add(cur.x+1, cur.y+1, cur.f+1)
				continue
			}
			clear := func(y int) bool {
				for _, wall := range walls {
					if wall[0] <= y && y < wall[1] {
						return true
					}
				}
				return false
			}
			if clear(cur.y - 1) {
				add(cur.x+1, cur.y-1, cur.f)
			}
			if clear(cur.y + 1) {
				add(cur.x+1, cur.y+1, cur.f+1)
			}
		}
	}
	return -1
}

func (p *parsed) String() string {
	var sb strings.Builder
	for x := range p.w + 1 {
		w, ok := p.m[x]
		if !ok {
			fmt.Fprintln(&sb, strings.Repeat(".", p.h))
			continue
		}
		fmt.Fprintf(&sb, "%s%s%s\n",
			strings.Repeat("#", w[0][0]),
			strings.Repeat(".", w[0][1]-w[0][0]),
			strings.Repeat("#", p.h-w[0][1]),
		)
	}
	return sb.String()
}
