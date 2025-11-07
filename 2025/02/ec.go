package main

import (
	"fmt"
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

func part2(a *C) int {
	return count(a, 10)
}

func part3(a *C) int {
	return count(a, 1)
}

func main() {
	r1 := part1(&C{145, 51})
	r2 := part2(&C{-4581, -68892})
	r3 := part3(&C{-4581, -68892})

	fmt.Printf("Part 1: %v\n", r1)
	fmt.Printf("Part 2: %v\n", r2)
	fmt.Printf("Part 3: %v\n", r3)
}
