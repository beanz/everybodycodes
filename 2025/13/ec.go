package main

import (
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := parse(i1)
	a2 := parse2(i2, 20252025)
	a3 := parse2(i3, 202520252025)
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

func parse(in []byte) int {
	nums := make([][]int, 2)
	nums[0] = append(nums[0], 1)
	j := 0
	for i := 0; i < len(in); i++ {
		var n int
		i, n = ChompUInt[int](in, i)
		nums[j] = append(nums[j], n)
		j = 1 - j
	}
	for i := len(nums[1]) - 1; i >= 0; i-- {
		nums[0] = append(nums[0], nums[1][i])
	}
	return nums[0][2025%len(nums[0])]
}

type r struct {
	s, e int
	dir  bool
}

func parse2(in []byte, target int) int {
	ranges := []r{{1, 2, true}}
	var rev []r
	dir := true
	size := 1
	for i := 0; i < len(in); i++ {
		var n, m int
		i, n = ChompUInt[int](in, i)
		i++
		i, m = ChompUInt[int](in, i)
		if dir {
			ranges = append(ranges, r{n, m + 1, true})
		} else {
			rev = append(rev, r{n, m + 1, false})
		}
		size += m + 1 - n
		dir = !dir
	}
	for i := len(rev) - 1; i >= 0; i-- {
		ranges = append(ranges, rev[i])
	}
	target %= size
	for _, r := range ranges {
		l := r.e - r.s
		if target >= l {
			target -= l
			continue
		}
		if r.dir {
			return r.s + target
		} else {
			return r.e - 1 - target
		}
	}
	return -1
}
