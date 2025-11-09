package main

import (
	"fmt"
	"math"
	"sort"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1, n1, m1 := 0, 0, math.MaxInt
	for i := 0; i < len(i1); i++ {
		var a int
		i, a = ChompUInt[int](i1, i)
		m1 = min(a, m1)
		a1 += a
		n1++
	}
	a1 -= n1 * m1

	a2, n2, m2 := 0, 0, math.MaxInt
	for i := 0; i < len(i2); i++ {
		var a int
		i, a = ChompUInt[int](i2, i)
		m2 = min(a, m2)
		a2 += a
		n2++
	}
	a2 -= n2 * m2

	var nums []int
	for i := 0; i < len(i3); i++ {
		var a int
		i, a = ChompUInt[int](i3, i)
		nums = append(nums, a)
	}
	sort.Ints(nums)
	m3 := nums[len(nums)/2]
	a3 := 0
	for _, a := range nums {
		if a < m3 {
			a3 += m3 - a
		} else {
			a3 += a - m3
		}
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
