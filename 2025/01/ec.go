package main

import (
	_ "embed"
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, string, string) {
	p1 := parse(i1)
	i := 0
	for _, m := range p1.moves {
		i = min(max(0, i+m), len(p1.names)-1)
	}
	a1 := p1.names[i]

	p2 := parse(i2)
	i = 0
	for _, m := range p2.moves {
		i = (len(p2.names) + m + i) % len(p2.names)
	}
	a2 := p2.names[i]

	p3 := parse(i3)
	for _, m := range p3.moves {
		m %= len(p3.names)
		if m < 0 {
			m += len(p3.names)
		}
		p3.names[0], p3.names[m] = p3.names[m], p3.names[0]
	}
	a3 := p3.names[0]
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	names []string
	moves []int
}

func parse(in []byte) *parsed {
	names := []string{}
	i := 0
	j := 0
OUTER:
	for ; i < len(in); i++ {
		switch in[i] {
		case '\n':
			names = append(names, string(in[j:i]))
			break OUTER
		case ',':
			names = append(names, string(in[j:i]))
			j = i + 1
		}
	}
	i += 2
	var moves []int
	for i < len(in) {
		m := 1
		if in[i] == 'L' {
			m = -1
		}
		j, n := ChompUInt[int](in, i+1)
		i = j + 1
		moves = append(moves, m*n)
	}
	return &parsed{names, moves}
}
