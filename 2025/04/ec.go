package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	num := []int{}
	for i := 0; i < len(i1); {
		j, n := ChompUInt[int](i1, i)
		i = j + 1
		num = append(num, n)
	}
	p1 := 2025 * num[0] / num[len(num)-1]

	num = num[:0]
	for i := 0; i < len(i2); {
		j, n := ChompUInt[int](i2, i)
		i = j + 1
		num = append(num, n)
	}
	p2 := int(math.Ceil(10000000000000.0 * float64(num[len(num)-1]) / float64(num[0])))

	i := 0
	j, s := ChompUInt[int](i3, i)
	i = j + 1
	for {
		j, n := ChompUInt[int](i3, i)
		i = j
		if i == len(i3) || i3[i] == '\n' {
			s = 100 * s / n
			break
		}
		j, m := ChompUInt[int](i3, i+1)
		i = j + 1
		s = s * m / n
	}
	p3 := s

	return p1, p2, p3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
