package main

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func part1(in []byte) int {
	p := parse(in)
	var a1 int
	n := 4
	if p.w < 16 {
		n = 3
	}
	todo := []*point{p.dragon}
	next := []*point{}
	for range n {
		for _, d := range todo {
			for _, m := range d.knightMoves() {
				c, keep := p.sheep1(m)
				if !keep {
					continue
				}
				next = append(next, m)
				a1 += c
			}
		}
		todo, next = next, todo
		next = next[:0]
	}
	return a1
}

func part2(in []byte) int {
	p := parse(in)
	var a int
	n := 20
	if p.w < 16 {
		n = 3
	}
	nextD := map[point]struct{}{}
	nextS := map[point]struct{}{}
	for range n {
		for d := range p.d {
			for _, m := range d.knightMoves() {
				c, keep := p.sheep2(m)
				if !keep {
					continue
				}
				nextD[*m] = struct{}{}
				a += c
			}
		}
		p.d, nextD = nextD, p.d
		clear(nextD)
		for s := range p.s {
			s.y++
			if s.y >= p.h {
				// survivor
				continue
			}
			i := s.x + s.y*p.w
			if p.m[i] == '.' {
				if _, ok := p.d[s]; ok {
					a++
				} else {
					nextS[s] = struct{}{}
				}
			} else {
				nextS[s] = struct{}{}
			}
		}
		p.s, nextS = nextS, p.s
		clear(nextS)
	}
	return a
}

type rec struct {
	d    point
	s    [8]byte
	n    int
	path string
}

func (r *rec) key() int {
	k := r.d.x
	k = (k << 4) + r.d.y
	for _, y := range r.s {
		k = (k << 4) + int(y&0xf)
	}
	return k
}

func part3(in []byte) int {
	p := parse(in)
	start := rec{*p.dragon, [8]byte{255, 255, 255, 255, 255, 255, 255, 255}, 0, ""}
	for k := range p.s {
		start.s[k.x] = byte(k.y)
		start.n++
	}
	home := bytes.Repeat([]byte{byte(p.h)}, 8)
	for x := 0; x < p.w-1; x++ {
		for y := p.h - 1; y > 0; y-- {
			if p.m[x+y*p.w] == '#' {
				home[x] = byte(y)
			} else {
				break
			}
		}
	}
	Lg(home)
	show := func(pos rec) string {
		var sb strings.Builder
		fmt.Fprintln(&sb, pos.path)
		fmt.Fprintf(&sb, "%d,%d %v %d\n", pos.d.x, pos.d.y, pos.s, pos.n)
		for y := range p.h {
			for x := range p.w - 1 {
				if p.m[x+y*p.w] == '#' {
					if pos.d.x == x && pos.d.y == y {
						fmt.Fprint(&sb, "d")
						continue
					}
					fmt.Fprint(&sb, "#")
					continue
				}
				if pos.d.x == x && pos.d.y == y {
					fmt.Fprint(&sb, "D")
					continue
				}
				if int(pos.s[x]) == y {
					fmt.Fprint(&sb, "S")
					continue
				}
				fmt.Fprint(&sb, ".")
			}
			fmt.Fprintln(&sb)
		}
		return sb.String()
	}
	chess := func(x, y byte) string {
		return string([]byte{'A' + x, '1' + y})
	}
	Lg(show(start))
	var moveDragon, moveSheep func(pos rec) int
	seenD := map[int]int{}
	seenS := map[int]int{}
	moveDragon = func(pos rec) int {
		if v, ok := seenD[pos.key()]; ok {
			return v
		}
		Lg("mD\n", show(pos))
		c := 0
		for _, m := range pos.d.knightMoves() {
			if !p.contains(m) {
				continue
			}
			ms := " D>" + chess(byte(m.x), byte(m.y))
			if int(pos.s[m.x]) == m.y {
				i := m.x + m.y*p.w
				if p.m[i] != '#' {
					if pos.n > 1 {
						ns := rec{*m, pos.s, pos.n - 1, pos.path + ms + "X"}
						ns.s[m.x] = 255
						c += moveSheep(ns)
					} else {
						Lg("all dead " + pos.path + ms + "X")
						c++
					}
					continue
				}
			}
			c += moveSheep(rec{*m, pos.s, pos.n, pos.path + ms})
		}
		seenD[pos.key()] = c
		return c
	}
	moveSheep = func(pos rec) int {
		if v, ok := seenS[pos.key()]; ok {
			return v
		}
		Lg("mS\n", show(pos))
		c := 0
		stuck := true
		for x, y := range pos.s {
			if y == 255 {
				continue
			}
			Lg("  ", x, y)
			if y == home[x]-1 {
				stuck = false
				continue
			}
			if pos.d.x == x && pos.d.y == int(y+1) {
				i := x + int(y+1)*p.w
				if p.m[i] != '#' {
					continue
				}
			}
			stuck = false
			ns := pos.s
			ns[x] = y + 1
			ms := " S>" + chess(byte(x), y+1)
			c += moveDragon(rec{pos.d, ns, pos.n, pos.path + ms})
		}
		seenS[pos.key()] = c
		if stuck {
			return moveDragon(pos)
		}
		return c
	}
	return moveSheep(start)
}

func parts(i1, i2, i3 []byte) (int, int, int) {
	return part1(i1), part2(i2), part3(i3)
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

func (p *point) moves() []*point {
	return []*point{
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x - 1, p.y},
	}
}

func (p *point) knightMoves() []*point {
	return []*point{
		{p.x + 1, p.y - 2},
		{p.x + 2, p.y - 1},
		{p.x + 2, p.y + 1},
		{p.x + 1, p.y + 2},
		{p.x - 1, p.y + 2},
		{p.x - 2, p.y + 1},
		{p.x - 2, p.y - 1},
		{p.x - 1, p.y - 2},
	}
}

type parsed struct {
	w, h   int
	dragon *point
	m      []byte
	s      map[point]struct{}
	d      map[point]struct{}
}

func parse(in []byte) *parsed {
	p := &parsed{m: in, s: map[point]struct{}{}, d: map[point]struct{}{}}
	p.w = bytes.IndexByte(in, '\n') + 1
	p.h = (1 + len(in)) / p.w
	for y := range p.h {
		for x := range p.w - 1 {
			i := x + y*p.w
			switch p.m[i] {
			case 'S':
				p.s[point{x, y}] = struct{}{}
				p.m[i] = '.'
			case 'D':
				p.dragon = &point{x, y}
				p.d[*p.dragon] = struct{}{}
				p.m[i] = '.'
			}
		}
	}
	return p
}

func (p *parsed) contains(xy *point) bool {
	if 0 <= xy.x && xy.x < p.w-1 && 0 <= xy.y && xy.y < p.h {
		return true
	}
	return false
}

func (p *parsed) sheep1(xy *point) (int, bool) {
	if !p.contains(xy) {
		return 0, false
	}
	i := xy.x + xy.y*p.w
	if _, ok := p.s[*xy]; ok {
		delete(p.s, *xy)
		p.m[i] = '#'
		return 1, true
	}
	switch p.m[i] {
	case '#':
		return 0, false
	case '.':
		p.m[i] = '#'
		return 0, true
	default:
		panic(fmt.Sprintf("unexpected state: %c", p.m[i]))
	}
}

func (p *parsed) sheep2(xy *point) (int, bool) {
	if !p.contains(xy) {
		return 0, false
	}
	i := xy.x + xy.y*p.w
	switch p.m[i] {
	case '#':
		return 0, true
	case '.':
		if _, ok := p.s[*xy]; ok {
			delete(p.s, *xy)
			return 1, true
		}
		return 0, true
	default:
		panic(fmt.Sprintf("unexpected state: %c", p.m[i]))
	}
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%dx%d %d,%d\n", p.w-1, p.h, p.dragon.x, p.dragon.y)
	fmt.Fprintf(&sb, "D=%d S=%d\n", len(p.d), len(p.s))
	for y := range p.h {
		for x := range p.w - 1 {
			if p.m[x+y*p.w] == '#' {
				fmt.Fprintf(&sb, "#")
				continue
			}
			xy := point{x, y}
			if _, ok := p.d[xy]; ok {
				fmt.Fprintf(&sb, "D")
				continue
			}
			if _, ok := p.s[xy]; ok {
				fmt.Fprintf(&sb, "S")
				continue
			}
			fmt.Fprintf(&sb, ".")
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
