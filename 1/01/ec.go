package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (uint, uint, uint) {
	a1 := solve(i1, eni)
	a2 := solve(i2, eni2)
	a3 := solve(i3, eni3)
	return a1, a2, a3
}

func solve(in []byte, fn func(a, exp, mod uint) uint) uint {
	var res uint
	for i := 0; i < len(in); i++ {
		var a, b, c, x, y, z, m uint
		i, a = ChompUInt[uint](in, i+2)
		i, b = ChompUInt[uint](in, i+3)
		i, c = ChompUInt[uint](in, i+3)
		i, x = ChompUInt[uint](in, i+3)
		i, y = ChompUInt[uint](in, i+3)
		i, z = ChompUInt[uint](in, i+3)
		i, m = ChompUInt[uint](in, i+3)
		r := fn(a, x, m) + fn(b, y, m) + fn(c, z, m)
		res = max(res, r)
	}
	return res
}

func eni3(n, exp, mod uint) uint {
	s := uint(1)
	t := uint(0)
	sums := []uint{0}
	seen := map[uint]uint{}
	var i uint
	for {
		s = (s * n) % mod
		t += s
		i++
		if p, ok := seen[s]; ok {
			cl := i - p
			cd := t - sums[p]
			rem := exp - i + 1
			q := rem / cl
			r := rem % cl
			return (t - s) + (q * cd) + sums[p+r-1] - sums[p-1]
		}
		seen[s] = uint(i)
		sums = append(sums, t)
	}
}

func eni2(n, exp, mod uint) uint {
	var r uint
	for j := uint(0); j < 5; j++ {
		a := ModExp(n, exp-j, mod)
		if a >= 100 {
			r = 1000*r + a
		} else if a >= 10 {
			r = 100*r + a
		} else {
			r = 10*r + a
		}
	}
	return r
}

func eni(n, exp, mod uint) uint {
	s := uint(1)
	var a uint
	m := uint(1)
	for range exp {
		s = (s * n) % mod
		a = a + s*m
		m *= 10
		if s >= 10 {
			m *= 10
		}
	}
	return a
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
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
