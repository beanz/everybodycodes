package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, int, [2]int) {
	p1 := parse(i1)
	var a1 string
	{
		for i := 0; i < 100; i++ {
			p1.iter()
		}
		a1 = p1.String()
	}
	p2 := parse(i2)
	p3 := parse(i3)
	return a1, p2.part2(), p3.part3()
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	n []int
	c [][][]byte
	i []int
}

func (p *parsed) part3() [2]int {
	pulls := 256
	cache := map[[2]int][2]int{}
	var sc func(o int, rem int) [2]int
	sc = func(o int, rem int) [2]int {
		if v, ok := cache[[2]int{o, rem}]; ok {
			return v
		}
		n := pulls - rem
		s := 0
		if n != 0 {
			s += p.score3(o, n)
			if rem == 0 {
				return [2]int{s, s}
			}
		}
		m := [2]int{math.MinInt, math.MaxInt}
		for i := -1; i <= 1; i++ {
			mnx := sc(o+i, rem-1)
			m[0], m[1] = max(m[0], mnx[0]), min(m[1], mnx[1])
		}
		v := [2]int{s + m[0], s + m[1]}
		cache[[2]int{o, rem}] = v
		return v
	}
	return sc(0, pulls)
}

func (p *parsed) part2() int {
	var a int
	target := 202420242024
	cycle := p.cycle()
	if cycle > target || cycle < 0 {
		cycle = 1819457640
	}
	for i := 1; i <= cycle; i++ {
		p.iter()
		sc := p.score()
		a += sc
	}
	n := target / cycle
	target -= cycle * n
	a *= n
	for i := 1; i <= target; i++ {
		p.iter()
		sc := p.score()
		a += sc
	}
	return a
}

func (p *parsed) iter() {
	for j := range p.i {
		p.i[j] += p.n[j]
		p.i[j] %= len(p.c[j])
	}
}

func (p *parsed) score3(o, n int) int {
	c := make(map[byte]int, 3*len(p.c))
	for j := 0; j < len(p.c); j++ {
		n := Mod(o+p.n[j]*n, len(p.c[j]))
		ct := p.c[j][n]
		c[ct[0]]++
		c[ct[2]]++
	}
	sc := 0
	for _, v := range c {
		if v >= 3 {
			sc += v - 2
		}
	}
	return sc
}

func (p *parsed) score() int {
	c := make(map[byte]int, 3*len(p.c))
	for j := 0; j < len(p.c); j++ {
		c[p.c[j][p.i[j]][0]]++
		c[p.c[j][p.i[j]][2]]++
	}
	sc := 0
	for _, v := range c {
		if v >= 3 {
			sc += v - 2
		}
	}
	return sc
}

func (p *parsed) cycle() int {
	c := 0
	for k := 0; k < len(p.c); k++ {
		n, cl := p.n[k], len(p.c[k])
		for n > cl {
			n -= cl
		}
		if c == 0 {
			c = cl
		} else {
			c = LCM(c, cl)
		}
	}
	return c
}

func (p *parsed) String() string {
	buf := make([]byte, len(p.c)*4-1)
	for j := 0; j < len(p.c)-1; j++ {
		buf[j*4+3] = ' '
	}
	for j := range p.c {
		copy(buf[j*4:j*4+3], p.c[j][p.i[j]])
	}
	return string(buf)
}

func parse(in []byte) *parsed {
	p := &parsed{}
	i := 0
	for ; i < len(in); i++ {
		j, n := ChompUInt[int](in, i)
		p.n = append(p.n, n)
		i = j
		if in[i] == '\n' {
			break
		}
	}
	i += 2
	j := 0
	p.c = make([][][]byte, len(p.n))
	for ; i < len(in); i++ {
		if in[i] == '\n' {
			j = 0
			continue
		}
		if in[i] != ' ' {
			p.c[j] = append(p.c[j], in[i:i+3])
		}
		i += 3
		j++
		if i < len(in) && in[i] == '\n' {
			j = 0
		}
	}
	p.i = make([]int, len(p.c))
	return p
}
