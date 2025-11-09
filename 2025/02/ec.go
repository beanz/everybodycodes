package main

import (
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

type C struct {
	x, y int
}

func (a *C) Add(b *C) *C {
	return &C{a.x + b.x, a.y + b.y}
}

func (a *C) Mul(b *C) *C {
	return &C{a.x*b.x - a.y*b.y, a.x*b.y + a.y*b.x}
}

func (a *C) Div(b *C) *C {
	return &C{a.x / b.x, a.y / b.y}
}

func (a *C) String() string {
	return fmt.Sprintf("[%d,%d]", a.x, a.y)
}

func parse(in []byte) *C {
	i, a := ChompInt[int](in, 3)
	_, b := ChompInt[int](in, i+1)
	return &C{a, b}
}

func part1(a *C) string {
	r := &C{0, 0}
	for i := 0; i < 3; i++ {
		r = r.Mul(r)
		r = r.Div(&C{10, 10})
		r = r.Add(a)
	}
	return r.String()
}

func engrave(a *C) bool {
	r := &C{0, 0}
	for i := 0; i < 100; i++ {
		r = r.Mul(r)
		r = r.Div(&C{100000, 100000})
		r = r.Add(a)
	}
	return -1000000 <= r.x && r.x <= 1000000 &&
		-1000000 <= r.y && r.y <= 1000000
}

func count(a *C, step int) int {
	c := 0
	for y := 0; y <= 1000; y += step {
		for x := 0; x <= 1000; x += step {
			a := &C{a.x + x, a.y + y}
			if engrave(a) {
				c++
			}
		}
	}
	return c
}

func parts(i1, i2, i3 []byte) (string, int, int) {
	a1 := parse(i1)
	a2 := parse(i2)
	a3 := parse(i3)
	return part1(a1), count(a2, 10), count(a3, 1)
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
