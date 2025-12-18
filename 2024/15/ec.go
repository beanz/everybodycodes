package main

import (
	"bytes"
	"fmt"
	"math/bits"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	return solve(i1), solve(i2), solve(i3)
}

func solve(in []byte) int {
	n := []int{0, -1, 0, 1, 0}
	w1 := bytes.IndexByte(in, '\n') + 1
	find := 0
	for ch := byte('A'); ch <= 'Z'; ch++ {
		if bytes.IndexByte(in, ch) != -1 {
			find |= (1 << (ch - 'A'))
		}
	}
	Lg(find)
	sx := bytes.IndexByte(in, '.')
	in[sx] = '^'
	seen := map[[3]int]struct{}{}
	best := bits.OnesCount(uint(find))
	todo := [][4]int{{sx, 1, find, 1}}
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		if _, ok := seen[[3]int{cur[0], cur[1], cur[2]}]; ok {
			continue
		}
		seen[[3]int{cur[0], cur[1], cur[2]}] = struct{}{}
		ch := in[cur[0]+cur[1]*w1]
		if ch == '^' {
			if cur[2] == 0 {
				return cur[3]
			}
			continue
		}
		bc := bits.OnesCount(uint(cur[2]))
		if bc > best+4 {
			continue
		}
		best = min(bc, best)
		if 'A' <= ch && ch <= 'Z' {
			//Lf("found %c at %v\n", ch, cur)
			bit := 1 << (ch - 'A')
			if cur[2]&bit != 0 {
				cur[2] ^= bit
			}
		}
		for i := 0; i+1 < len(n); i++ {
			nx, ny, ns := cur[0]+n[i], cur[1]+n[i+1], cur[3]+1
			if in[nx+ny*w1] == '#' || in[nx+ny*w1] == '~' {
				continue
			}
			todo = append(todo, [4]int{nx, ny, cur[2], ns})
		}
	}
	return -1
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
