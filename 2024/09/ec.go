package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	var a1 int
	{
		cache := make([]*int, 20000)
		for i := 0; i < len(i1); i++ {
			j, n := ChompUInt[int](i1, i)
			i = j
			b := beetles(n, []int{10, 5, 3, 1}, cache)
			a1 += b
		}
	}
	var a2 int
	{
		cache := make([]*int, 20000)
		for i := 0; i < len(i2); i++ {
			j, n := ChompUInt[int](i2, i)
			i = j
			b := beetles(n, []int{30, 25, 24, 20, 16, 15, 10, 5, 3, 1}, cache)
			a2 += b
		}
	}
	var a3 int
	{
		cache := make([]*int, 500000)
		for i := 0; i < len(i3); i++ {
			j, n := ChompUInt[int](i3, i)
			i = j
			minimun := math.MaxInt
			for k := -50; k < 50; k++ {
				n0 := n/2 - k
				n1 := n - n0
				b0 := beetles(n0, []int{101, 100, 75, 74, 50, 49, 38, 37, 30, 25, 24, 20, 16, 15, 10, 5, 3, 1}, cache)
				b1 := beetles(n1, []int{101, 100, 75, 74, 50, 49, 38, 37, 30, 25, 24, 20, 16, 15, 10, 5, 3, 1}, cache)
				minimun = min(minimun, b0+b1)
			}
			a3 += minimun
		}
	}
	return a1, a2, a3
}

func beetles(target int, v []int, cache []*int) int {
	var aux func(n int) int
	aux = func(n int) int {
		if n == 0 {
			return 0
		}
		if r := cache[n]; r != nil {
			return *r
		}
		minimum := math.MaxInt
		for _, e := range v {
			if n >= e {
				minimum = min(minimum, 1+aux(n-e))
			}
		}
		cache[n] = &minimum
		return minimum
	}
	return aux(target)
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
