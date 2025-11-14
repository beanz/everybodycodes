package main

import (
	"bytes"
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.run(4, map[byte]int{'A': 1})

	p2 := parse(i2)
	a2 := p2.run(10, map[byte]int{'Z': 1})

	p3 := parse2(i3)
	var a3 int
	{
		mn, mx := math.MaxInt, math.MinInt
		for k := range p3.r {
			n := p3.run(20, map[byte]int{k: 1})
			if n < mn {
				mn = n
			}
			if n > mx {
				mx = n
			}
		}
		a3 = mx - mn
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	r map[byte][]byte
}

func (p *parsed) run(n int, start map[byte]int) int {
	res := 0
	m := start
	next := map[byte]int{}
	for range n {
		for k, v := range m {
			for _, to := range p.r[k] {
				next[to] += v
			}
		}
		next, m = m, next
		clear(next)
	}
	for _, v := range m {
		res += v
	}
	return res
}

func parse(in []byte) *parsed {
	p := &parsed{map[byte][]byte{}}
	for line := range bytes.SplitSeq(in, []byte{'\n'}) {
		from, rest, _ := bytes.Cut(line, []byte{':'})
		for to := range bytes.SplitSeq(rest, []byte{','}) {
			p.r[from[0]] = append(p.r[from[0]], to[0])
		}
	}
	return p
}

func parse2(in []byte) *parsed {
	m := NewIdentifierMap[string, byte]()
	p := &parsed{map[byte][]byte{}}
	for line := range bytes.SplitSeq(in, []byte{'\n'}) {
		from, rest, _ := bytes.Cut(line, []byte{':'})
		fromID := m.Add(string(from))
		for to := range bytes.SplitSeq(rest, []byte{','}) {
			toID := m.Add(string(to))
			p.r[fromID] = append(p.r[fromID], toID)
		}
	}
	return p
}
