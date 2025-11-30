package main

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.solve1()
	p2 := parse(i2)
	a2 := p2.solve2()
	p3 := parse(i3)
	a3 := p3.solve3()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	in   []byte
	w, h int
}

func parse(in []byte) *parsed {
	p := &parsed{in: in}
	p.w = 1 + bytes.IndexByte(in, '\n')
	p.h = (1 + len(in)) / p.w
	return p
}

func (p *parsed) solve1() int {
	a := 0
	for y := range p.h - 1 {
		for x := range p.w - 2 {
			if p.in[x+y*p.w] != 'T' {
				continue
			}
			if p.in[x+1+y*p.w] == 'T' {
				a++
			}
			if (x+y)%2 == 1 && p.in[x+(y+1)*p.w] == 'T' {
				a++
			}
		}
	}
	return a
}

func (p *parsed) solve2() int {
	i := bytes.IndexByte(p.in, 'S')
	type rec struct {
		x, y, s int
	}
	seen := make([]bool, len(p.in))
	sx, sy := i%p.w, i/p.w
	todo := []rec{{sx, sy, 0}}
	seen[sx+sy*p.w] = true
	add := func(x, y, s int) {
		if !(0 <= x && x < p.w-1 && 0 <= y && y < p.h) {
			return
		}
		if p.in[x+y*p.w] != 'T' && p.in[x+y*p.w] != 'E' {
			return
		}
		if seen[x+y*p.w] {
			return
		}
		seen[x+y*p.w] = true
		todo = append(todo, rec{x, y, s})
	}
	for len(todo) > 0 {
		cur := todo[0]
		Dg(cur)
		todo = todo[1:]
		if p.in[cur.x+cur.y*p.w] == 'E' {
			return cur.s
		}
		if (cur.x+cur.y)%2 == 0 {
			add(cur.x, cur.y-1, cur.s+1)
		} else {
			add(cur.x, cur.y+1, cur.s+1)
		}
		add(cur.x+1, cur.y, cur.s+1)
		add(cur.x-1, cur.y, cur.s+1)
	}
	return -1
}

func (p *parsed) solve3() int {
	i := bytes.IndexByte(p.in, 'S')
	type rec struct {
		x, y, s int
	}
	l := len(p.in)
	seen := make([]bool, 3*l)
	sx, sy := i%p.w, i/p.w
	todo := []rec{{sx, sy, 0}}
	seen[sx+sy*p.w] = true
	add := func(x, y, s int) {
		if !(0 <= x && x < p.w-1 && 0 <= y && y < p.h) {
			return
		}
		if p.in[x+y*p.w] != 'T' && p.in[x+y*p.w] != 'E' {
			return
		}
		k := x + y*p.w + (s%3)*l
		if seen[k] {
			return
		}
		seen[k] = true
		todo = append(todo, rec{x, y, s})
	}
	wD2 := (p.w - 1) / 2
	for len(todo) > 0 {
		cur := todo[0]
		Dg(cur)
		todo = todo[1:]
		if p.in[cur.x+cur.y*p.w] == 'E' {
			return cur.s
		}
		cur.x, cur.y = wD2-(cur.x-cur.y)/2+cur.y, wD2-(cur.x-cur.y+1)/2-cur.y
		add(cur.x, cur.y, cur.s+1)
		if (cur.x+cur.y)%2 == 0 {
			add(cur.x, cur.y-1, cur.s+1)
		} else {
			add(cur.x, cur.y+1, cur.s+1)
		}
		add(cur.x+1, cur.y, cur.s+1)
		add(cur.x-1, cur.y, cur.s+1)
	}
	return -1
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
