package main

import (
	"bytes"
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

var fb = []byte{'R', 'G', 'B'}

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := 1
	{
		for i := 0; i < len(i1); i++ {
			if i1[i] == '\n' {
				break
			}
			bolt := fb[(a1-1)%3]
			if i1[i] != bolt {
				Df("%d %c %s\n", a1, bolt, string(i1))
				if i+1 == len(i1) {
					i1[i] = '-'
					break
				}
				a1++
			}
			i1[i] = '-'
		}
	}
	a2 := solve(i2, 100)
	a3 := solve(i3, 100000)
	return a1, a2, a3
}

func solve(in []byte, rep int) int {
	a := 0
	buf := bytes.Repeat(in, rep)
	h := len(buf) / 2
	l, r := buf[:h], buf[h:]
	var cur byte
	for len(l) > 0 {
		cur, l = l[0], l[1:]
		bolt := fb[a%3]
		a++
		if cur == bolt {
			if (len(l)+len(r))%2 != 0 {
				r = r[1:]
			}
		} else if len(r) > len(l) {
			cur, r = r[0], r[1:]
			l = append(l, cur)
		}
	}
	return a
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
