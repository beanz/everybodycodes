package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	var a1 int
	{
		nails := 32
		if len(i1) < 30 {
			nails = 8
		}
		j, prev := ChompUInt[int](i1, 0)
		for i := j + 1; i < len(i1); i++ {
			j, n := ChompUInt[int](i1, i)
			i = j
			if Abs(prev-n) == nails/2 {
				a1++
			}
			prev = n
		}
	}

	var a2 int
	{
		j, prev := ChompUInt[int](i2, 0)
		var partitions []partition
		for i := j + 1; i < len(i2); i++ {
			j, n := ChompUInt[int](i2, i)
			i = j
			a, b := prev, n
			if a > b {
				a, b = b, a
			}
			a2 += cuts(partitions, a, b)
			partitions = append(partitions, partition{a, b})
			prev = n
		}
	}

	var a3 int
	{
		j, prev := ChompUInt[int](i3, 0)
		var partitions []partition
		for i := j + 1; i < len(i3); i++ {
			j, n := ChompUInt[int](i3, i)
			i = j
			a, b := prev, n
			if a > b {
				a, b = b, a
			}
			a3 += cuts(partitions, a, b)
			partitions = append(partitions, partition{a, b})
			prev = n
		}
		nails := 256
		if len(i3) < 30 {
			nails = 8
		}
		a3 = math.MinInt
		for i := 1; i <= nails; i++ {
			for j := i + 1; j <= nails; j++ {
				a3 = max(a3, cuts(partitions, i, j))
			}
		}
	}

	return a1, a2, a3
}

type partition struct {
	mn, mx int
}

func cuts(partitions []partition, a, b int) (res int) {
	for _, p := range partitions {
		aEq := a == p.mn || a == p.mx
		bEq := b == p.mn || b == p.mx
		if aEq && bEq {
			res++
			continue
		}
		if aEq || bEq {
			continue
		}
		aIn := p.mn <= a && a <= p.mx
		bIn := p.mn <= b && b <= p.mx
		if aIn != bIn {
			res++
		}
	}
	return
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
