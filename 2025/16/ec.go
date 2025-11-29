package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	var a1 int
	{
		factors := []int{}
		for i := 0; i < len(i1); i++ {
			var n int
			i, n = ChompUInt[int](i1, i)
			factors = append(factors, n)
		}
		for _, f := range factors {
			for i := 1; i <= 90; i++ {
				if i%f == 0 {
					a1++
				}
			}
		}
	}
	var a2 int
	{
		wall := []int{}
		for i := 0; i < len(i2); i++ {
			var n int
			i, n = ChompUInt[int](i2, i)
			wall = append(wall, n)
		}
		a2 = 1
		for i := range wall {
			n := wall[i]
			if n == 0 {
				continue
			}
			a2 *= (i + 1)
			for j := i + 1; j < len(wall); j++ {
				if (j+1)%(i+1) == 0 {
					wall[j]--
				}
			}
		}
	}
	var a3 int
	{
		var wall []int
		for i := 0; i < len(i3); i++ {
			var n int
			i, n = ChompUInt[int](i3, i)
			wall = append(wall, n)
		}
		var factors []int
		for i := range wall {
			n := wall[i]
			if n == 0 {
				continue
			}
			factors = append(factors, i+1)
			for j := i + 1; j < len(wall); j++ {
				if (j+1)%(i+1) == 0 {
					wall[j]--
				}
			}
		}
		blocks := func(l int) int {
			b := l
			for _, f := range factors[1:] {
				b += l / f
			}
			return b
		}
		limit := 202520252025000
		var l int
		for l = 1; blocks(l) < limit; l *= 2 {
		}
		h := l
		l /= 2
		for l < h-1 {
			m := (l + h) / 2
			if blocks(m) > limit {
				h = m
			} else {
				l = m
			}
		}
		a3 = l
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
}

func parse(in []byte) *parsed {
	p := &parsed{}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
