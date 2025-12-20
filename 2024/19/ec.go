package main

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, string, string) {
	p1 := parse(i1)
	p1.fastRound()
	p2 := parse(i2)
	for range 100 {
		p2.fastRound()
	}
	p3 := parse(i3)
	for i := 0; i < 100000000; i++ {
		p3.fastRound()
		w := p3.word()
		if w != "" {
			Dg(i, w)
			break
		}
	}
	return p1.word(), p2.word(), p3.word()
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	r    []byte
	m    []byte
	a    []byte
	p    []int
	w, h int
}

func parse(in []byte) *parsed {
	p := &parsed{}
	i := 0
	for ; in[i] != '\n'; i++ {
	}
	p.r = in[:i]
	i += 2
	p.m = in[i:]
	for ; in[i] != '\n'; i++ {
		p.w++
	}
	p.h = (1 + len(p.m)) / (p.w + 1)
	p.p = make([]int, len(p.m))
	for i := 0; i < len(p.p); i++ {
		p.p[i] = i
	}
	p.permRound()
	p.a = make([]byte, len(p.m))
	return p
}

func (p *parsed) word() string {
	s := bytes.IndexByte(p.m, '>')
	e := bytes.IndexByte(p.m, '<')
	if s > e {
		return ""
	}
	sx, sy := s%(p.w+1), s/(p.w+1)
	ex, ey := e%(p.w+1), e/(p.w+1)
	if sy != ey {
		return ""
	}
	w := string(p.m[s+1 : e])
	if w == "WIN" || w == "VICTORY" {
		return w
	}
	if ex-sx != 17 {
		return ""
	}
	Df("%q\n", w)
	for _, c := range w {
		if '0' <= c && c <= '9' {
			continue
		}
		return ""
	}
	return w
}

func (p *parsed) permRound() {
	i := 0
	for y := 1; y < p.h-1; y++ {
		for x := 1; x < p.w-1; x++ {
			r := p.r[i%len(p.r)]
			if r == 'R' {
				p.permRotateCW(x, y)
			} else {
				p.permRotateCCW(x, y)
			}
			i++
		}
	}
}

func (p *parsed) permRotateCW(x, y int) {
	i := func(x, y int) int {
		return x + y*(p.w+1)
	}
	p.p[i(x-1, y-1)],
		p.p[i(x+0, y-1)],
		p.p[i(x+1, y-1)],
		p.p[i(x+1, y+0)],
		p.p[i(x+1, y+1)],
		p.p[i(x+0, y+1)],
		p.p[i(x-1, y+1)],
		p.p[i(x-1, y+0)] =
		p.p[i(x-1, y+0)],
		p.p[i(x-1, y-1)],
		p.p[i(x+0, y-1)],
		p.p[i(x+1, y-1)],
		p.p[i(x+1, y+0)],
		p.p[i(x+1, y+1)],
		p.p[i(x+0, y+1)],
		p.p[i(x-1, y+1)]
}

func (p *parsed) permRotateCCW(x, y int) {
	i := func(x, y int) int {
		return x + y*(p.w+1)
	}
	p.p[i(x-1, y-1)],
		p.p[i(x+0, y-1)],
		p.p[i(x+1, y-1)],
		p.p[i(x+1, y+0)],
		p.p[i(x+1, y+1)],
		p.p[i(x+0, y+1)],
		p.p[i(x-1, y+1)],
		p.p[i(x-1, y+0)] =
		p.p[i(x+0, y-1)],
		p.p[i(x+1, y-1)],
		p.p[i(x+1, y+0)],
		p.p[i(x+1, y+1)],
		p.p[i(x+0, y+1)],
		p.p[i(x-1, y+1)],
		p.p[i(x-1, y+0)],
		p.p[i(x-1, y-1)]
}

func (p *parsed) fastRound() {
	for i := 0; i < len(p.m); i++ {
		p.a[i] = p.m[p.p[i]]
	}
	p.a, p.m = p.m, p.a
}

func (p *parsed) round() {
	i := 0
	for y := 1; y < p.h-1; y++ {
		for x := 1; x < p.w-1; x++ {
			r := p.r[i%len(p.r)]
			if r == 'R' {
				p.rotateCW(x, y)
			} else {
				p.rotateCCW(x, y)
			}
			i++
		}
	}
}

func (p *parsed) rotateCW(x, y int) {
	i := func(x, y int) int {
		return x + y*(p.w+1)
	}
	p.m[i(x-1, y-1)],
		p.m[i(x+0, y-1)],
		p.m[i(x+1, y-1)],
		p.m[i(x+1, y+0)],
		p.m[i(x+1, y+1)],
		p.m[i(x+0, y+1)],
		p.m[i(x-1, y+1)],
		p.m[i(x-1, y+0)] =
		p.m[i(x-1, y+0)],
		p.m[i(x-1, y-1)],
		p.m[i(x+0, y-1)],
		p.m[i(x+1, y-1)],
		p.m[i(x+1, y+0)],
		p.m[i(x+1, y+1)],
		p.m[i(x+0, y+1)],
		p.m[i(x-1, y+1)]
}

func (p *parsed) rotateCCW(x, y int) {
	i := func(x, y int) int {
		return x + y*(p.w+1)
	}
	p.m[i(x-1, y-1)],
		p.m[i(x+0, y-1)],
		p.m[i(x+1, y-1)],
		p.m[i(x+1, y+0)],
		p.m[i(x+1, y+1)],
		p.m[i(x+0, y+1)],
		p.m[i(x-1, y+1)],
		p.m[i(x-1, y+0)] =
		p.m[i(x+0, y-1)],
		p.m[i(x+1, y-1)],
		p.m[i(x+1, y+0)],
		p.m[i(x+1, y+1)],
		p.m[i(x+0, y+1)],
		p.m[i(x-1, y+1)],
		p.m[i(x-1, y+0)],
		p.m[i(x-1, y-1)]
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "R: %d,%d %q\n", p.w, p.h, string(p.r))
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			fmt.Fprintf(&sb, "%c", p.m[x+y*(p.w+1)])
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
