package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func main() {
	i1 := Input(1, "")
	num := []int{}
	for i := 0; i < len(i1); {
		j, n := ChompUInt[int](i1, i)
		i = j + 1
		num = append(num, n)
	}
	p1 := 2025 * num[0] / num[len(num)-1]
	i2 := Input(2, "")
	num = num[:0]
	for i := 0; i < len(i2); {
		j, n := ChompUInt[int](i2, i)
		i = j + 1
		num = append(num, n)
	}
	p2 := int(math.Ceil(10000000000000.0 * float64(num[len(num)-1]) / float64(num[0])))
	i3 := Input(3, "")
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

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Part 3: %d\n", p3)
}
