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
	a1 := p1.part1()
	p2 := parse(i2)
	a2 := p2.part1()
	p3 := parse(i3)
	a3 := p3.part3()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	m    []byte
	w, h int
	p    int
}

func (p *parsed) part3() int {
	o := []int{0, -1, 0, 1, 0}
	count := make([]int, len(p.m))
	for i := bytes.IndexByte(p.m, 'P'); i != -1; i = bytes.IndexByte(p.m, 'P') {
		sx, sy := i%p.w, i/p.w
		p.m[i] = 'p'
		seen := make([]bool, len(p.m))
		todo := [][3]int{{sx, sy, 0}}
		for len(todo) > 0 {
			cur := todo[0]
			todo = todo[1:]
			i := cur[0] + cur[1]*p.w
			if seen[i] {
				continue
			}
			seen[i] = true
			count[i] += cur[2]
			for i := 0; i < 4; i++ {
				nx, ny := cur[0]+o[i], cur[1]+o[i+1]
				ch := p.m[nx+ny*p.w]
				if ch == '.' || ch == 'P' || ch == 'p' {
					todo = append(todo, [3]int{nx, ny, cur[2] + 1})
				}
			}
		}
	}
	a := math.MaxInt
	for i, v := range count {
		if v == 0 || p.m[i] != '.' {
			continue
		}
		a = min(a, v)
	}
	return a
}

func (p *parsed) part1() int {
	o := []int{0, -1, 0, 1, 0}
	p.m[1*p.w] = '#'
	todo := [][3]int{{0, 1, 0}}
	if p.m[(p.w-2)+(p.h-2)*p.w] == '.' {
		todo = append(todo, [3]int{p.w - 2, p.h - 2, 0})
		p.m[(p.w-2)+(p.h-2)*p.w] = '#'
	}
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		ch := p.m[cur[0]+cur[1]*p.w]
		if ch == 'P' {
			p.p--
			if p.p == 0 {
				return cur[2]
			}
		}
		if ch == '~' {
			continue
		}
		p.m[cur[0]+cur[1]*p.w] = '~'
		for i := 0; i < 4; i++ {
			nx, ny := cur[0]+o[i], cur[1]+o[i+1]
			ch := p.m[nx+ny*p.w]
			if ch == '.' || ch == 'P' {
				todo = append(todo, [3]int{nx, ny, cur[2] + 1})
			}
		}
	}
	return -1
}

func parse(in []byte) *parsed {
	w := bytes.IndexByte(in, '\n') + 1
	h := (1 + len(in)) / w
	p := &parsed{m: in, w: w, h: h, p: bytes.Count(in, []byte{'P'})}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
