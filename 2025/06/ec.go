package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := 0
	c := make([]int, 3)
	for _, ch := range i1 {
		if ch == '\n' {
			continue
		}
		if ch != 'A' && ch != 'a' {
			continue
		}
		if ch&32 == 0 {
			c[int(ch-'A')]++
			continue
		}
		a1 += c[int(ch-'a')]
	}
	a2 := 0
	{
		c := make([]int, 3)
		for _, ch := range i2 {
			if ch == '\n' {
				continue
			}
			if ch&32 == 0 {
				c[int(ch-'A')]++
				continue
			}
			a2 += c[int(ch-'a')]
		}
	}
	a3 := 0
	{
		repeats, dist := 1000, 1000
		if len(i3) < 50 {
			repeats, dist = 2, 10
		}
		get := func(i int) byte {
			return i3[(len(i3)+i)%len(i3)]
		}
		for i := range len(i3) {
			ch := get(i)
			if ch != 'a' && ch != 'b' && ch != 'c' {
				continue
			}
			c := 0
			mch := ch - 32
			for j := i - dist; j <= i+dist; j++ {
				och := get(j)
				if och == mch {
					c += repeats
					if (j < 0 && len(i3)-1-j >= len(i3)) || j >= len(i3) {
						c--
					}
				}
			}
			a3 += c
		}

	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
}

func parse(in []byte) *parsed {
	return &parsed{}
}

func (p *parsed) String() string {
	var sb strings.Builder
	return sb.String()
}
