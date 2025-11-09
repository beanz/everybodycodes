package main

import (
	"fmt"
	"sort"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	num := []int{}
	for i := 0; i < len(i1); {
		j, n := ChompUInt[int](i1, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p1 := num[0]
	p := num[0]
	for i := 1; i < len(num); i++ {
		if p == num[i] {
			continue
		}
		p = num[i]
		p1 += p
	}

	num = num[:0]
	for i := 0; i < len(i2); {
		j, n := ChompUInt[int](i2, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p2 := num[0]
	p = num[0]
	c := 1
	for i := 1; i < len(num); i++ {
		if p == num[i] {
			continue
		}
		c++
		p = num[i]
		p2 += p
		if c == 20 {
			break
		}
	}

	num = num[:0]
	for i := 0; i < len(i3); {
		j, n := ChompUInt[int](i3, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p3 := 0
	for i := 0; i < len(num); {
		n := num[i]
		c := 0
		for ; i < len(num); i++ {
			if num[i] == n {
				c++
			} else {
				break
			}
		}
		if p3 < c {
			p3 = c
		}
	}
	return p1, p2, p3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
