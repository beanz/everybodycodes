package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, string, int) {
	p1 := parse(i1)
	a1 := p1.part1(10000)
	p2 := parse(i2)
	a2 := p2.part2()
	p3 := parse(i3)
	a3 := p3.part3()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type (
	parsed struct {
		dice   []*die
		target []byte
		w, h   int
	}
	die struct {
		i     int
		faces []int
		pulse int
		seed  int
		rn    int
	}
)

func (d *die) roll() int {
	d.rn++
	s := d.rn * d.pulse
	d.i += s
	d.pulse += s
	d.pulse %= d.seed
	d.pulse += 1 + d.rn + d.seed
	return d.faces[d.i%len(d.faces)]
}

func (p *parsed) part3() int {
	m := map[int]struct{}{}
	l := make([]byte, 10240)
	type rec struct {
		x, y int
		i    int
	}
	for _, d := range p.dice {
		for i := range l {
			l[i] = '0' + byte(d.roll())
		}
		todo := []rec{}
		for y := 0; y < p.h; y++ {
			for x := 0; x < p.w-1; x++ {
				i := x + y*p.w
				if p.target[i] == l[0] {
					todo = append(todo, rec{x, y, 0})
				}
			}
		}
		add := func(x, y, i int) {
			if !(0 <= x && x < p.w-1 && 0 <= y && y < p.h) {
				return
			}
			if p.target[x+y*p.w] != l[i] {
				return
			}
			todo = append(todo, rec{x, y, i})
		}
		seen := map[rec]struct{}{}
		for len(todo) > 0 {
			cur := todo[0]
			todo = todo[1:]
			i := cur.x + cur.y*p.w
			if cur.i >= 10240 {
				continue
			}
			m[i] = struct{}{}
			if _, ok := seen[cur]; ok {
				continue
			}
			seen[cur] = struct{}{}
			add(cur.x, cur.y, cur.i+1)
			add(cur.x, cur.y-1, cur.i+1)
			add(cur.x+1, cur.y, cur.i+1)
			add(cur.x, cur.y+1, cur.i+1)
			add(cur.x-1, cur.y, cur.i+1)
		}
	}
	return len(m)
}

func (p *parsed) part2() string {
	score := make([][2]int, 0, len(p.dice))
	for i, d := range p.dice {
		j := 0
		c := 1
		for ; ; c++ {
			r := d.roll()
			if byte('0'+r) == p.target[j] {
				j++
				if j == len(p.target) {
					break
				}
			}
		}
		score = append(score, [2]int{i + 1, c})
	}
	slices.SortFunc(score, func(a, b [2]int) int {
		return a[1] - b[1]
	})
	r := make([]string, len(score))
	for i, e := range score {
		r[i] = fmt.Sprint(e[0])
	}
	return strings.Join(r, ",")
}

func (p *parsed) part1(n int) int {
	s := 0
	for r := 1; ; r++ {
		s += p.roll()
		if s > n {
			return r
		}
	}
}

func (p *parsed) roll() int {
	s := 0
	for _, d := range p.dice {
		s += d.roll()
	}
	return s
}

func parse(in []byte) *parsed {
	p := &parsed{}
	for i := 0; i < len(in); i++ {
		if in[i] == '\n' {
			p.target = in[i+1:]
			p.w = bytes.IndexByte(p.target, '\n')
			if p.w != -1 {
				p.w++
				p.h = (1 + len(p.target)) / p.w
			}
			break
		}
		j, _ := ChompUInt[int](in, i)
		i = j + 9
		d := &die{}
		for ; i < len(in); i++ {
			j, n := ChompInt[int](in, i)
			d.faces = append(d.faces, n)
			i = j
			if in[j] == ']' {
				break
			}
		}
		i, d.seed = ChompInt[int](in, i+7)
		d.pulse = d.seed
		p.dice = append(p.dice, d)
	}
	return p
}
